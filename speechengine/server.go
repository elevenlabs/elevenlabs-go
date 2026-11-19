// Hand-written (listed in .fernignore): the inbound WebSocket server the ElevenLabs Speech Engine
// API dials into, mirroring SpeechEngineServer and attach() in the Python, TypeScript, PHP, .NET
// and Rust SDKs.

package speechengine

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/coder/websocket"
)

const defaultServerAddr = ":3001"

// How long a client has to send its request headers, so a peer that opens a socket and never
// completes the upgrade cannot hold it open indefinitely.
const readHeaderTimeout = 10 * time.Second

// SpeechEngineServerOptions configures SpeechEngineServer and SpeechEngineHandler.
type SpeechEngineServerOptions struct {
	// Addr is the address SpeechEngineServer listens on. Defaults to ":3001". Ignored by
	// SpeechEngineHandler.
	Addr string

	// Path, when set, rejects connections to any other request path with 404.
	Path string

	// APIKey verifies inbound connections. Defaults to the ELEVENLABS_API_KEY environment
	// variable. SpeechEngineResource fills it from the client it came from.
	APIKey string

	// DisableAuth accepts connections without verifying their JWT.
	//
	// Insecure: only enable it when the server is protected another way, typically by an IP
	// allowlist restricted to ElevenLabs' egress ranges. Otherwise anyone who can reach the server
	// can open a session and spend your compute and LLM quota.
	DisableAuth bool

	// Debug logs connection and protocol events with the "[SpeechEngine]" prefix.
	Debug bool

	// Callbacks are run for every session.
	Callbacks SpeechEngineCallbacks
}

// SpeechEngineHandler is an http.Handler that accepts inbound Speech Engine connections: it
// verifies the X-Elevenlabs-Speech-Engine-Authorization JWT, upgrades the request to a WebSocket
// and runs a session until the connection ends. Mount it on an existing server, or use
// SpeechEngineServer for a standalone one.
type SpeechEngineHandler struct {
	options SpeechEngineServerOptions
	apiKey  string

	mu       sync.Mutex
	closed   bool
	sessions map[*sessionEntry]struct{}
	active   sync.WaitGroup
}

type sessionEntry struct {
	cancel context.CancelFunc
}

// NewSpeechEngineHandler creates a handler. Without an API key (and without DisableAuth) every
// connection is rejected with 401.
func NewSpeechEngineHandler(opts SpeechEngineServerOptions) *SpeechEngineHandler {
	apiKey := opts.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("ELEVENLABS_API_KEY")
	}
	if opts.DisableAuth {
		log.Print("[SpeechEngine] authentication is disabled: incoming connections will NOT be verified. Make sure the server is protected by an IP allowlist restricting traffic to ElevenLabs.")
	} else if apiKey == "" {
		log.Print("[SpeechEngine] no API key configured: every connection will be rejected. Set SpeechEngineServerOptions.APIKey or ELEVENLABS_API_KEY.")
	}
	return &SpeechEngineHandler{
		options:  opts,
		apiKey:   apiKey,
		sessions: make(map[*sessionEntry]struct{}),
	}
}

func (h *SpeechEngineHandler) logf(format string, args ...any) {
	if h.options.Debug {
		log.Printf("[SpeechEngine] "+format, args...)
	}
}

// ServeHTTP implements http.Handler.
func (h *SpeechEngineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.options.Path != "" && r.URL.Path != h.options.Path {
		h.logf("path mismatch: expected %s, got %s", h.options.Path, r.URL.Path)
		http.NotFound(w, r)
		return
	}
	if !h.options.DisableAuth {
		if err := verifyRequestHeader(r.Header, h.apiKey); err != nil {
			h.logf("rejected connection: %v", err)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	entry := &sessionEntry{cancel: cancel}
	h.mu.Lock()
	if h.closed {
		h.mu.Unlock()
		http.Error(w, "speech engine server is shutting down", http.StatusServiceUnavailable)
		return
	}
	h.sessions[entry] = struct{}{}
	h.active.Add(1)
	h.mu.Unlock()
	defer func() {
		h.mu.Lock()
		delete(h.sessions, entry)
		h.mu.Unlock()
		h.active.Done()
	}()

	// Origin checks protect browsers from cross-site WebSocket hijacking; the caller here is the
	// ElevenLabs API, which is authenticated by the JWT above instead.
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		h.logf("websocket upgrade failed: %v", err)
		return
	}
	h.logf("connection established, creating session")
	if err := ServeSpeechEngineConn(ctx, conn, h.options.Callbacks, &SpeechEngineSessionOptions{Debug: h.options.Debug}); err != nil {
		h.logf("session ended with error: %v", err)
	}
}

// Close stops accepting connections and closes every active session, waiting for them to end
// or for ctx to be done. It does not affect the server the handler is mounted on.
func (h *SpeechEngineHandler) Close(ctx context.Context) error {
	h.mu.Lock()
	h.closed = true
	for entry := range h.sessions {
		entry.cancel()
	}
	h.mu.Unlock()

	done := make(chan struct{})
	go func() {
		h.active.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SpeechEngineServer is a standalone server for inbound Speech Engine connections. To share a
// port with an existing net/http server, mount a SpeechEngineHandler instead.
type SpeechEngineServer struct {
	options SpeechEngineServerOptions
	handler *SpeechEngineHandler
	server  *http.Server

	mu       sync.Mutex
	listener net.Listener
}

// NewSpeechEngineServer creates a server. Call ListenAndServe or Serve to start it.
func NewSpeechEngineServer(opts SpeechEngineServerOptions) *SpeechEngineServer {
	if opts.APIKey == "" {
		opts.APIKey = os.Getenv("ELEVENLABS_API_KEY")
	}
	if opts.Addr == "" {
		opts.Addr = defaultServerAddr
	}
	handler := NewSpeechEngineHandler(opts)
	return &SpeechEngineServer{
		options: opts,
		handler: handler,
		server: &http.Server{
			Addr:              opts.Addr,
			Handler:           handler,
			ReadHeaderTimeout: readHeaderTimeout,
		},
	}
}

func (s *SpeechEngineServer) checkAuth() error {
	if s.options.APIKey == "" && !s.options.DisableAuth {
		return errors.New(`speech engine: SpeechEngineServer requires an API key to verify incoming connections. Set SpeechEngineServerOptions.APIKey or the ELEVENLABS_API_KEY environment variable. To run without authentication, set DisableAuth — but only behind an IP allowlist`)
	}
	return nil
}

// ListenAndServe listens on the configured address and serves connections until Shutdown is
// called, then returns http.ErrServerClosed.
func (s *SpeechEngineServer) ListenAndServe() error {
	if err := s.checkAuth(); err != nil {
		return err
	}
	listener, err := net.Listen("tcp", s.options.Addr)
	if err != nil {
		return err
	}
	return s.Serve(listener)
}

// Serve serves connections accepted on listener until Shutdown is called, then returns
// http.ErrServerClosed.
func (s *SpeechEngineServer) Serve(listener net.Listener) error {
	if err := s.checkAuth(); err != nil {
		listener.Close()
		return err
	}
	s.mu.Lock()
	s.listener = listener
	s.mu.Unlock()
	s.handler.logf("speech engine server listening on %s", listener.Addr())
	return s.server.Serve(listener)
}

// Addr returns the address the server is listening on, or nil before it starts. Useful when
// listening on port 0.
func (s *SpeechEngineServer) Addr() net.Addr {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listener == nil {
		return nil
	}
	return s.listener.Addr()
}

// Shutdown stops accepting connections, closes every active session and waits for them to end
// or for ctx to be done.
func (s *SpeechEngineServer) Shutdown(ctx context.Context) error {
	// http.Server.Shutdown does not track hijacked (WebSocket) connections, so the handler closes
	// its sessions itself.
	err := s.server.Shutdown(ctx)
	if closeErr := s.handler.Close(ctx); err == nil {
		err = closeErr
	}
	return err
}
