// Hand-written (listed in .fernignore).

package speechengine

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// fakeTransport records outbound frames.
type fakeTransport struct {
	mu     sync.Mutex
	frames []map[string]any
	sent   chan struct{}
	closed atomic.Bool
}

func newFakeTransport() *fakeTransport {
	return &fakeTransport{sent: make(chan struct{}, 1024)}
}

func (f *fakeTransport) WriteText(_ context.Context, data []byte) error {
	var frame map[string]any
	if err := json.Unmarshal(data, &frame); err != nil {
		return err
	}
	f.mu.Lock()
	f.frames = append(f.frames, frame)
	f.mu.Unlock()
	f.sent <- struct{}{}
	return nil
}

func (f *fakeTransport) Close() error {
	f.closed.Store(true)
	return nil
}

func (f *fakeTransport) snapshot() []map[string]any {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]map[string]any(nil), f.frames...)
}

// waitFrames waits until at least n frames have been written.
func (f *fakeTransport) waitFrames(t *testing.T, n int) []map[string]any {
	t.Helper()
	deadline := time.After(5 * time.Second)
	for len(f.snapshot()) < n {
		select {
		case <-f.sent:
		case <-deadline:
			t.Fatalf("timed out waiting for %d frames, got %v", n, f.snapshot())
		}
	}
	return f.snapshot()
}

func transcriptFrame(eventID int, content string) []byte {
	data, _ := json.Marshal(map[string]any{
		"type":            "user_transcript",
		"event_id":        eventID,
		"user_transcript": []map[string]string{{"role": "user", "content": content}},
	})
	return data
}

func TestPingIsAnsweredWithPong(t *testing.T) {
	transport := newFakeTransport()
	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{}, nil)
	session.HandleMessage([]byte(`{"type":"ping"}`))
	frames := transport.waitFrames(t, 1)
	if frames[0]["type"] != "pong" {
		t.Fatalf("expected pong, got %v", frames[0])
	}
}

func TestInitSetsConversationID(t *testing.T) {
	var seen string
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnInit: func(id string, _ *SpeechEngineSession) { seen = id },
	}, nil)
	session.HandleMessage([]byte(`{"type":"init","conversation_id":"conv_42"}`))
	if session.ConversationID() != "conv_42" || seen != "conv_42" {
		t.Fatalf("conversation id not set: session=%q callback=%q", session.ConversationID(), seen)
	}
}

func TestTranscriptHandlerSendsAResponse(t *testing.T) {
	transport := newFakeTransport()
	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{
		OnTranscript: func(ctx context.Context, transcript []ConversationMessage, s *SpeechEngineSession) error {
			if len(transcript) != 1 || transcript[0].Role != "user" || transcript[0].Content != "hello" {
				t.Errorf("unexpected transcript: %v", transcript)
			}
			return s.SendResponse(ctx, "hi there")
		},
	}, nil)
	session.HandleMessage(transcriptFrame(7, "hello"))
	session.Wait()
	frames := transport.waitFrames(t, 2)
	if frames[0]["type"] != "agent_response" || frames[0]["content"] != "hi there" || frames[0]["event_id"] != float64(7) || frames[0]["is_final"] != false {
		t.Fatalf("unexpected first frame: %v", frames[0])
	}
	if frames[1]["content"] != "" || frames[1]["is_final"] != true || frames[1]["event_id"] != float64(7) {
		t.Fatalf("unexpected final frame: %v", frames[1])
	}
}

func TestTranscriptHandlerStreamsMixedChunks(t *testing.T) {
	transport := newFakeTransport()
	type chatChunk struct {
		Choices []struct {
			Delta struct {
				Content string `json:"content"`
			} `json:"delta"`
		} `json:"choices"`
	}
	var structChunk chatChunk
	structChunk.Choices = make([]struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	}, 1)
	structChunk.Choices[0].Delta.Content = "chat "

	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{
		OnTranscript: func(ctx context.Context, _ []ConversationMessage, s *SpeechEngineSession) error {
			chunks := make(chan any, 8)
			chunks <- "plain "
			chunks <- json.RawMessage(`{"type":"response.output_text.delta","delta":"openai "}`)
			chunks <- structChunk
			chunks <- map[string]any{"type": "content_block_delta", "delta": map[string]any{"type": "text_delta", "text": "anthropic "}}
			chunks <- []byte(`{"candidates":[{"content":{"parts":[{"thought":true,"text":"hmm"},{"text":"gemini"}]}}]}`)
			chunks <- map[string]any{"type": "message_start"} // no text: skipped
			close(chunks)
			return s.SendResponseStream(ctx, chunks)
		},
	}, nil)
	session.HandleMessage(transcriptFrame(3, "go"))
	session.Wait()
	frames := transport.waitFrames(t, 6)
	var contents []string
	for _, f := range frames[:5] {
		contents = append(contents, f["content"].(string))
		if f["is_final"] != false {
			t.Fatalf("content frame marked final: %v", f)
		}
	}
	want := []string{"plain ", "openai ", "chat ", "anthropic ", "gemini"}
	for i := range want {
		if contents[i] != want[i] {
			t.Fatalf("contents = %q, want %q", contents, want)
		}
	}
	if len(frames) != 6 || frames[5]["is_final"] != true {
		t.Fatalf("expected a trailing final frame, got %v", frames)
	}
}

func TestSendResponseOutsideAHandlerIsANoop(t *testing.T) {
	transport := newFakeTransport()
	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{}, nil)
	if err := session.SendResponse(context.Background(), "nope"); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if frames := transport.snapshot(); len(frames) != 0 {
		t.Fatalf("expected no frames, got %v", frames)
	}
}

func TestNewTranscriptInterruptsTheRunningHandler(t *testing.T) {
	transport := newFakeTransport()
	started := make(chan struct{})
	interrupted := make(chan error, 1)
	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{
		OnTranscript: func(ctx context.Context, transcript []ConversationMessage, s *SpeechEngineSession) error {
			if transcript[0].Content == "first" {
				close(started)
				<-ctx.Done()
				// A response for the interrupted turn is dropped.
				interrupted <- s.SendResponse(ctx, "stale")
				return ctx.Err()
			}
			return s.SendResponse(ctx, "fresh")
		},
	}, nil)

	session.HandleMessage(transcriptFrame(1, "first"))
	<-started
	session.HandleMessage(transcriptFrame(2, "second"))
	session.Wait()

	if err := <-interrupted; !errors.Is(err, context.Canceled) {
		t.Fatalf("interrupted SendResponse = %v, want context.Canceled", err)
	}
	frames := transport.waitFrames(t, 2)
	for _, f := range frames {
		if f["event_id"] != float64(2) {
			t.Fatalf("only the newer turn may respond, got %v", frames)
		}
	}
	if frames[0]["content"] != "fresh" {
		t.Fatalf("unexpected frames: %v", frames)
	}
}

func TestDuplicateTranscriptIsSkipped(t *testing.T) {
	var calls atomic.Int32
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnTranscript: func(context.Context, []ConversationMessage, *SpeechEngineSession) error {
			calls.Add(1)
			return nil
		},
	}, nil)
	session.HandleMessage(transcriptFrame(5, "a"))
	session.HandleMessage(transcriptFrame(5, "a"))
	session.Wait()
	if calls.Load() != 1 {
		t.Fatalf("expected 1 handler call, got %d", calls.Load())
	}
}

func TestHandlerErrorIsReported(t *testing.T) {
	errCh := make(chan error, 1)
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnTranscript: func(context.Context, []ConversationMessage, *SpeechEngineSession) error {
			return errors.New("llm failed")
		},
		OnError: func(err error, _ *SpeechEngineSession) { errCh <- err },
	}, nil)
	session.HandleMessage(transcriptFrame(1, "a"))
	session.Wait()
	select {
	case err := <-errCh:
		if err.Error() != "llm failed" {
			t.Fatalf("unexpected error: %v", err)
		}
	default:
		t.Fatal("handler error was not reported")
	}
}

func TestCloseMessageFiresOnCloseAndClosesTheTransport(t *testing.T) {
	transport := newFakeTransport()
	var fired bool
	session := NewSpeechEngineSession(transport, SpeechEngineCallbacks{
		OnClose: func(s *SpeechEngineSession) { fired = s.IsOpen() },
	}, nil)
	session.HandleMessage([]byte(`{"type":"close"}`))
	if !fired {
		t.Fatal("OnClose must run while the session is still open")
	}
	if session.IsOpen() || !transport.closed.Load() {
		t.Fatal("session and transport must be closed after a close message")
	}
}

func TestCloseCancelsARunningHandler(t *testing.T) {
	started := make(chan struct{})
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnTranscript: func(ctx context.Context, _ []ConversationMessage, s *SpeechEngineSession) error {
			close(started)
			<-ctx.Done()
			if err := s.SendResponse(ctx, "late"); !errors.Is(err, ErrSessionClosed) {
				t.Errorf("SendResponse after close = %v, want ErrSessionClosed", err)
			}
			return nil
		},
	}, nil)
	session.HandleMessage(transcriptFrame(1, "a"))
	<-started
	session.Close()
	session.Close() // idempotent
	session.Wait()
}

func TestDisconnectFiresOnceEvenAfterClose(t *testing.T) {
	var count atomic.Int32
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnDisconnect: func(*SpeechEngineSession) { count.Add(1) },
	}, nil)
	session.HandleMessage([]byte(`{"type":"close"}`))
	session.HandleDisconnect()
	session.HandleDisconnect()
	if count.Load() != 1 {
		t.Fatalf("expected OnDisconnect once, got %d", count.Load())
	}
}

func TestMalformedFramesAndErrorMessagesAreReported(t *testing.T) {
	var errs []string
	session := NewSpeechEngineSession(newFakeTransport(), SpeechEngineCallbacks{
		OnError: func(err error, _ *SpeechEngineSession) { errs = append(errs, err.Error()) },
	}, nil)
	session.HandleMessage([]byte("not json"))
	session.HandleMessage([]byte(`{"type":"error","message":"upstream failed"}`))
	session.HandleMessage([]byte(`{"type":"something_new"}`)) // ignored
	if len(errs) != 2 || errs[1] != "upstream failed" {
		t.Fatalf("unexpected errors: %q", errs)
	}
}

func TestExtractText(t *testing.T) {
	cases := []struct {
		chunk any
		text  string
		ok    bool
	}{
		{"x", "x", true},
		{map[string]any{"choices": []any{map[string]any{"delta": map[string]any{"content": "c"}}}}, "c", true},
		{json.RawMessage(`{"candidates":[{"content":{"parts":[{"thought":true,"text":"r"},{"text":"answer"}]}}]}`), "answer", true},
		{map[string]any{"unknown": "shape"}, "", false},
		{nil, "", false},
		{42, "", false},
	}
	for _, tc := range cases {
		text, ok := ExtractText(tc.chunk)
		if text != tc.text || ok != tc.ok {
			t.Errorf("ExtractText(%v) = %q, %v; want %q, %v", tc.chunk, text, ok, tc.text, tc.ok)
		}
	}
}

// Cancelling the context a connection is served under cancels the running turn at once, without
// waiting for the read loop to see the cancellation (it can be blocked writing a pong).
func TestCancellingTheServeContextCancelsTheRunningTurn(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	stopped := make(chan struct{})
	session := newSession(parent, newFakeTransport(), SpeechEngineCallbacks{
		OnTranscript: func(ctx context.Context, _ []ConversationMessage, _ *SpeechEngineSession) error {
			close(started)
			<-ctx.Done()
			close(stopped)
			return ctx.Err()
		},
	}, nil)
	session.HandleMessage([]byte(`{"type":"user_transcript","event_id":1,"user_transcript":[{"role":"user","content":"hi"}]}`))
	<-started
	cancel()
	select {
	case <-stopped:
	case <-time.After(5 * time.Second):
		t.Fatal("the turn's context must be cancelled with the serve context")
	}
	session.Wait()
}
