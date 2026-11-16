package texttodialoguemulticontext

import (
	context "context"
	json "encoding/json"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	http "net/http"
	url "net/url"
	sync "sync"
)

// TextToDialogueMultiContextConnectRequest holds the handshake parameters of the text_to_dialogue_multi_context channel.
type TextToDialogueMultiContextConnectRequest struct {
	ModelId                *string
	OutputFormat           *string
	LanguageCode           *string
	SyncAlignment          *bool
	ApplyTextNormalization *types.TextToDialogueMultiContextRealtimeApplyTextNormalization
	Seed                   *int
	EnableLogging          *bool
}

// Realtime opens a connection to the text_to_dialogue_multi_context channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *TextToDialogueMultiContextConnectRequest, opts ...option.RequestOption) (*TextToDialogueMultiContextSocket, error) {
	if request == nil {
		request = &TextToDialogueMultiContextConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "model_id", request.ModelId)
	internal.AddWebSocketQuery(query, "output_format", request.OutputFormat)
	internal.AddWebSocketQuery(query, "language_code", request.LanguageCode)
	internal.AddWebSocketQuery(query, "sync_alignment", request.SyncAlignment)
	internal.AddWebSocketQuery(query, "apply_text_normalization", request.ApplyTextNormalization)
	internal.AddWebSocketQuery(query, "seed", request.Seed)
	internal.AddWebSocketQuery(query, "enable_logging", request.EnableLogging)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/text-to-dialogue/multi-stream-input"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &TextToDialogueMultiContextSocket{}
	connection, err := internal.ConnectWebSocket(ctx, internal.WebSocketConnectionParams{
		URL: address,
		Header: func() (http.Header, error) {
			return c.options.HandshakeHeader(nil, options)
		},
		HTTPClient: internal.WebSocketHTTPClient(options, c.options),
		Reconnect:  internal.WebSocketReconnect(c.options, options),
		Events:     socket.events(),
	})
	if err != nil {
		return nil, err
	}
	socket.connection = connection
	return socket, nil
}

// TextToDialogueMultiContextSocket is an open connection to the text_to_dialogue_multi_context channel.
//
// Stream expressive dialogue audio for multiple independent streams (contexts) multiplexed over a single WebSocket connection.
//
// Each context, identified by a client-chosen `context_id`, behaves like an independent [Text to Dialogue WebSocket](/docs/api-reference/text-to-dialogue/ttd-websocket) session: it registers its own voices and settings, buffers its own text, and produces its own audio stream. This is useful for scenarios requiring concurrent or interleaved dialogue generations, such as conversational AI applications that need to handle interruptions.
//
// The connection uses Eleven v3 and v4 dialogue models (`model_id` must start with `eleven_v3` or `eleven_v4`). The default model is `eleven_v3_conversational`.
//
// ## Context setup
// - Every message **must** include a `context_id`. A message containing only `close_socket` is the exception.
// - The first message for a new `context_id` creates that context and **must** include `voices` (voice IDs to register for the context). Optional `voice_settings` and `pronunciation_dictionary_locators` are only accepted on this first message.
// - For `eleven_v3_conversational` and `eleven_v4` models, only **one** voice ID may be registered per context. For `eleven_v3`, you may register up to **10** voices per context.
// - A connection can hold at most **5** simultaneous contexts; close a context to free a slot.
//
// ## Streaming text
// - Send `inputs`: an array of `{ "text", "voice_id", "new_turn"? }`. Each `voice_id` must be registered for that context. Text for the same turn is buffered per context until the server has enough context, then partial audio chunks tagged with the `context_id` are emitted.
// - Set `new_turn` to `true` (or switch `voice_id`) to finalize the current prosody segment and start a new speaker turn.
//
// ## Control messages
// - `flush`: force generation of the context's buffered text.
// - `close_context`: flush the context's remaining audio, emit its `is_final` message, and close it. Other contexts stay open.
// - `close_socket`: flush and close **all** contexts, then close the connection.
// - `keep_alive`: reset the context's **20 second** inactivity timeout (no generation). A context idle for longer is automatically flushed and closed (its `is_final` message is sent); other contexts are unaffected.
//
// Protocol errors — a missing `context_id`, an unregistered voice, messaging a context that is closing, or exceeding the context limit — send an error payload and close the whole connection.
//
// ## Authentication
// Authentication is connection-level, not per context: use the `xi-api-key` or `Authorization` header, `single_use_token` query parameter, or include `xi_api_key`, `authorization`, or `single_use_token` in the first message of the connection. Anonymous sessions are rejected.
//
// For a single dialogue stream per connection, see the [Text to Dialogue WebSocket](/docs/api-reference/text-to-dialogue/ttd-websocket). For non-streaming dialogue over HTTP, see [Create dialogue](/docs/api-reference/text-to-dialogue/convert) and [Stream dialogue](/docs/api-reference/text-to-dialogue/stream).
type TextToDialogueMultiContextSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   textToDialogueMultiContextSocketHandlers
}

type textToDialogueMultiContextSocketHandlers struct {
	onReceiveTextToDialogueWebsocketMessageMulti func(*types.ReceiveTextToDialogueWebsocketMessageMulti)
	onRaw                                        func(string)
	onUnknown                                    func(json.RawMessage)
	onError                                      func(error)
	onConnected                                  func()
	onDisconnected                               func(core.DisconnectReason)
}

// OnReceiveTextToDialogueWebsocketMessageMulti registers the handler for ReceiveTextToDialogueWebsocketMessageMulti messages, replacing any earlier one.
func (s *TextToDialogueMultiContextSocket) OnReceiveTextToDialogueWebsocketMessageMulti(handler func(*types.ReceiveTextToDialogueWebsocketMessageMulti)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveTextToDialogueWebsocketMessageMulti = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *TextToDialogueMultiContextSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *TextToDialogueMultiContextSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *TextToDialogueMultiContextSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *TextToDialogueMultiContextSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *TextToDialogueMultiContextSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendTextToDialogueMultiClientMessage sends TextToDialogueMultiClientMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *TextToDialogueMultiContextSocket) SendTextToDialogueMultiClientMessage(ctx context.Context, message *types.TextToDialogueMultiClientMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *TextToDialogueMultiContextSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *TextToDialogueMultiContextSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *TextToDialogueMultiContextSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *TextToDialogueMultiContextSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *TextToDialogueMultiContextSocket) Close() error {
	return s.connection.Close()
}

func (s *TextToDialogueMultiContextSocket) snapshot() textToDialogueMultiContextSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *TextToDialogueMultiContextSocket) events() internal.WebSocketEvents {
	return internal.WebSocketEvents{
		OnText:   s.receiveText,
		OnBinary: s.receiveBinary,
		OnConnected: func() {
			if handler := s.snapshot().onConnected; handler != nil {
				handler()
			}
		},
		OnDisconnected: func(reason core.DisconnectReason) {
			if handler := s.snapshot().onDisconnected; handler != nil {
				handler(reason)
			}
		},
		OnError: s.reportError,
	}
}

func (s *TextToDialogueMultiContextSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *TextToDialogueMultiContextSocket) receiveText(data []byte) {
	handlers := s.snapshot()
	if handlers.onRaw != nil {
		handlers.onRaw(string(data))
	}
	frame, err := internal.ParseWebSocketFrame(data)
	if err != nil {
		s.reportError(err)
		return
	}
	{
		var message types.ReceiveTextToDialogueWebsocketMessageMulti
		if frame.Decode(&message) {
			if handlers.onReceiveTextToDialogueWebsocketMessageMulti != nil {
				handlers.onReceiveTextToDialogueWebsocketMessageMulti(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *TextToDialogueMultiContextSocket) receiveBinary(_ []byte) {}
