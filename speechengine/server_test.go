// Hand-written (listed in .fernignore). End-to-end tests: a real server on a local port and a
// coder/websocket client standing in for the ElevenLabs API.

package speechengine

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
)

func dialEngine(t *testing.T, ctx context.Context, url string, token string) (*websocket.Conn, *http.Response, error) {
	t.Helper()
	header := http.Header{}
	if token != "" {
		header.Set(AuthorizationHeader, token)
	}
	return websocket.Dial(ctx, url, &websocket.DialOptions{HTTPHeader: header})
}

func readFrame(t *testing.T, ctx context.Context, conn *websocket.Conn) map[string]any {
	t.Helper()
	_, data, err := conn.Read(ctx)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var frame map[string]any
	if err := json.Unmarshal(data, &frame); err != nil {
		t.Fatalf("decode %q: %v", data, err)
	}
	return frame
}

func writeFrame(t *testing.T, ctx context.Context, conn *websocket.Conn, frame string) {
	t.Helper()
	if err := conn.Write(ctx, websocket.MessageText, []byte(frame)); err != nil {
		t.Fatalf("write: %v", err)
	}
}

// echoCallbacks reply to each transcript with "you said: <last message>" and report lifecycle
// events on the returned channels.
func echoCallbacks() (SpeechEngineCallbacks, chan string, chan struct{}) {
	initCh := make(chan string, 1)
	disconnected := make(chan struct{}, 1)
	return SpeechEngineCallbacks{
		OnInit: func(id string, _ *SpeechEngineSession) { initCh <- id },
		OnTranscript: func(ctx context.Context, transcript []ConversationMessage, s *SpeechEngineSession) error {
			return s.SendResponse(ctx, "you said: "+transcript[len(transcript)-1].Content)
		},
		OnDisconnect: func(*SpeechEngineSession) { disconnected <- struct{}{} },
	}, initCh, disconnected
}

func startServer(t *testing.T, opts SpeechEngineServerOptions) (*SpeechEngineServer, string) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	server := NewSpeechEngineServer(opts)
	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
		if err := <-served; !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("Serve returned %v, want http.ErrServerClosed", err)
		}
	})
	return server, "ws://" + listener.Addr().String()
}

func TestServerEndToEnd(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callbacks, initCh, disconnected := echoCallbacks()
	_, url := startServer(t, SpeechEngineServerOptions{APIKey: goldenAPIKey, Path: "/ws", Callbacks: callbacks})

	conn, _, err := dialEngine(t, ctx, url+"/ws", "Bearer "+goldenToken)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.CloseNow()

	writeFrame(t, ctx, conn, `{"type":"init","conversation_id":"conv_1"}`)
	if id := <-initCh; id != "conv_1" {
		t.Fatalf("init conversation id = %q", id)
	}

	writeFrame(t, ctx, conn, `{"type":"ping"}`)
	if frame := readFrame(t, ctx, conn); frame["type"] != "pong" {
		t.Fatalf("expected pong, got %v", frame)
	}

	writeFrame(t, ctx, conn, `{"type":"user_transcript","event_id":9,"user_transcript":[{"role":"user","content":"hello"}]}`)
	first := readFrame(t, ctx, conn)
	final := readFrame(t, ctx, conn)
	if first["content"] != "you said: hello" || first["event_id"] != float64(9) || first["is_final"] != false {
		t.Fatalf("unexpected response: %v", first)
	}
	if final["is_final"] != true || final["event_id"] != float64(9) {
		t.Fatalf("unexpected final frame: %v", final)
	}

	// A close message ends the conversation: the server closes the connection.
	writeFrame(t, ctx, conn, `{"type":"close"}`)
	if _, _, err := conn.Read(ctx); websocket.CloseStatus(err) != websocket.StatusNormalClosure {
		t.Fatalf("expected a normal closure, got %v", err)
	}
	select {
	case <-disconnected:
	case <-ctx.Done():
		t.Fatal("OnDisconnect was not called")
	}
}

func TestServerRejectsUnverifiedConnections(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, url := startServer(t, SpeechEngineServerOptions{APIKey: goldenAPIKey, Path: "/ws"})

	cases := []struct {
		name   string
		path   string
		token  string
		status int
	}{
		{"missing header", "/ws", "", http.StatusUnauthorized},
		{"wrong key", "/ws", strings.Replace(goldenToken, "UObM", "XXXX", 1), http.StatusUnauthorized},
		{"wrong path", "/other", goldenToken, http.StatusNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, resp, err := dialEngine(t, ctx, url+tc.path, tc.token)
			if err == nil {
				conn.CloseNow()
				t.Fatal("connection must be rejected")
			}
			if resp == nil || resp.StatusCode != tc.status {
				t.Fatalf("expected HTTP %d, got %v (%v)", tc.status, resp, err)
			}
		})
	}
}

func TestServerRequiresAnAPIKey(t *testing.T) {
	t.Setenv("ELEVENLABS_API_KEY", "")
	err := NewSpeechEngineServer(SpeechEngineServerOptions{Addr: "127.0.0.1:0"}).ListenAndServe()
	if err == nil || !strings.Contains(err.Error(), "requires an API key") {
		t.Fatalf("expected a missing-key error, got %v", err)
	}
}

func TestDisableAuthAcceptsUnsignedConnections(t *testing.T) {
	t.Setenv("ELEVENLABS_API_KEY", "")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, url := startServer(t, SpeechEngineServerOptions{DisableAuth: true})
	conn, _, err := dialEngine(t, ctx, url, "")
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	conn.Close(websocket.StatusNormalClosure, "")
}

func TestShutdownClosesActiveSessions(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callbacks, initCh, disconnected := echoCallbacks()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	server := NewSpeechEngineServer(SpeechEngineServerOptions{APIKey: goldenAPIKey, Callbacks: callbacks})
	go server.Serve(listener)

	conn, _, err := dialEngine(t, ctx, "ws://"+listener.Addr().String(), goldenToken)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer conn.CloseNow()
	writeFrame(t, ctx, conn, `{"type":"init","conversation_id":"conv_1"}`)
	<-initCh

	if err := server.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}
	select {
	case <-disconnected:
	case <-ctx.Done():
		t.Fatal("Shutdown must close active sessions")
	}
	if _, _, err := conn.Read(ctx); err == nil {
		t.Fatal("the client connection must be closed")
	}
}

func TestAttachToAnExistingServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callbacks, _, _ := echoCallbacks()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("ok")) })
	engine := NewSpeechEngineResource("seng_1", goldenAPIKey, nil)
	handler := engine.Attach(mux, "/api/speech-engine/ws", SpeechEngineServerOptions{Callbacks: callbacks})
	httpServer := httptest.NewServer(mux)
	defer httpServer.Close()
	url := "ws" + strings.TrimPrefix(httpServer.URL, "http")

	// The rest of the app keeps working on the same server.
	resp, err := http.Get(httpServer.URL + "/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("health check: %v %v", resp, err)
	}
	resp.Body.Close()

	conn, _, err := dialEngine(t, ctx, url+"/api/speech-engine/ws", goldenToken)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	writeFrame(t, ctx, conn, `{"type":"user_transcript","event_id":1,"user_transcript":[{"role":"user","content":"hi"}]}`)
	if frame := readFrame(t, ctx, conn); frame["content"] != "you said: hi" {
		t.Fatalf("unexpected response: %v", frame)
	}
	readFrame(t, ctx, conn) // final frame

	// Closing the handler ends its sessions and rejects new connections, without stopping the
	// HTTP server.
	if err := handler.Close(ctx); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, _, err := conn.Read(ctx); err == nil {
		t.Fatal("the session must be closed")
	}
	if _, resp, err := dialEngine(t, ctx, url+"/api/speech-engine/ws", goldenToken); err == nil || resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 after Close, got %v %v", resp, err)
	}
	if resp, err := http.Get(httpServer.URL + "/health"); err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("health check after Close: %v %v", resp, err)
	} else {
		resp.Body.Close()
	}
}

func TestResourceServeStopsWithItsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := listener.Addr().String()
	listener.Close()

	engine := NewSpeechEngineResource("seng_1", goldenAPIKey, nil)
	done := make(chan error, 1)
	go func() { done <- engine.Serve(ctx, SpeechEngineServerOptions{Addr: addr}) }()

	dialCtx, dialCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dialCancel()
	var conn *websocket.Conn
	for conn == nil {
		conn, _, err = dialEngine(t, dialCtx, "ws://"+addr, goldenToken)
		if err != nil {
			if dialCtx.Err() != nil {
				t.Fatalf("server never came up: %v", err)
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
	defer conn.CloseNow()

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Serve returned %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Serve did not return after its context was cancelled")
	}
}
