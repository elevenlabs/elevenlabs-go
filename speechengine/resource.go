// Hand-written (listed in .fernignore): the handle returned by the speech engine
// Create/Get/Update/Duplicate methods, mirroring SpeechEngineResource in the Python, TypeScript,
// PHP, .NET and Rust SDKs.

package speechengine

import (
	"context"
	"errors"
	"net/http"
	"time"

	types "github.com/elevenlabs/elevenlabs-go/types"
)

// How long Serve waits for active sessions to end once its context is done.
const serveShutdownTimeout = 5 * time.Second

// SpeechEngineResource is a speech engine, with the server-side helpers for accepting the
// connections the ElevenLabs Speech Engine API opens to your server. It carries the API key of
// the client it came from, which verifies those connections.
//
// Use Serve for a standalone server, Handler or Attach to share an existing net/http server, or
// VerifyRequest and ServeSpeechEngineConn to manage the WebSocket upgrade yourself.
type SpeechEngineResource struct {
	// EngineID is the speech engine ID.
	EngineID string
	// Config is the API response the resource was built from, or nil when it was built without
	// an API call (SpeechEngineClient.Resource, NewSpeechEngineResource).
	Config *types.SpeechEngineResponse

	apiKey string
}

// NewSpeechEngineResource builds a resource without an API call. apiKey verifies inbound
// connections; config may be nil.
func NewSpeechEngineResource(engineID string, apiKey string, config *types.SpeechEngineResponse) *SpeechEngineResource {
	return &SpeechEngineResource{EngineID: engineID, Config: config, apiKey: apiKey}
}

// VerifyRequest checks that an inbound request carries a valid
// X-Elevenlabs-Speech-Engine-Authorization JWT, returning nil when it does and an error
// describing the problem otherwise. Only needed when you handle the WebSocket upgrade yourself;
// Serve, Handler and Attach verify automatically.
func (r *SpeechEngineResource) VerifyRequest(req *http.Request) error {
	if req == nil {
		return errors.New("speech engine: nil request")
	}
	return verifyRequestHeader(req.Header, r.apiKey)
}

func (r *SpeechEngineResource) withAPIKey(opts SpeechEngineServerOptions) SpeechEngineServerOptions {
	if opts.APIKey == "" {
		opts.APIKey = r.apiKey
	}
	return opts
}

// Handler returns an http.Handler that accepts this engine's connections, for mounting on an
// existing server. opts.APIKey defaults to the resource's key; opts.Addr is ignored.
func (r *SpeechEngineResource) Handler(opts SpeechEngineServerOptions) *SpeechEngineHandler {
	return NewSpeechEngineHandler(r.withAPIKey(opts))
}

// HandleRegistrar is satisfied by *http.ServeMux and by most routers (chi, gorilla/mux).
type HandleRegistrar interface {
	Handle(pattern string, handler http.Handler)
}

// Attach mounts a handler for this engine's connections on mux at path, and returns it so the
// caller can Close it later. Routing is left to mux: the handler doesn't check the request path
// again, because a subrouter (chi's Route/Mount, gorilla's PathPrefix().Subrouter()) registers a
// relative pattern while r.URL.Path keeps the full path. opts.Path is ignored.
func (r *SpeechEngineResource) Attach(mux HandleRegistrar, path string, opts SpeechEngineServerOptions) *SpeechEngineHandler {
	opts.Path = ""
	handler := r.Handler(opts)
	mux.Handle(path, handler)
	return handler
}

// Server returns a standalone server for this engine. opts.APIKey defaults to the resource's key.
func (r *SpeechEngineResource) Server(opts SpeechEngineServerOptions) *SpeechEngineServer {
	return NewSpeechEngineServer(r.withAPIKey(opts))
}

// Serve runs a standalone server for this engine on opts.Addr (default ":3001") until ctx is
// done, then shuts it down and returns nil. It returns early with an error if the server cannot
// start.
func (r *SpeechEngineResource) Serve(ctx context.Context, opts SpeechEngineServerOptions) error {
	server := r.Server(opts)
	errCh := make(chan error, 1)
	go func() { errCh <- server.ListenAndServe() }()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), serveShutdownTimeout)
	defer cancel()
	err := server.Shutdown(shutdownCtx)
	if serveErr := <-errCh; serveErr != nil && !errors.Is(serveErr, http.ErrServerClosed) && err == nil {
		err = serveErr
	}
	return err
}
