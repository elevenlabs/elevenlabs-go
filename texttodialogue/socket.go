package texttodialogue

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

// TextToDialogueConnectRequest holds the handshake parameters of the text_to_dialogue channel.
type TextToDialogueConnectRequest struct {
	ModelId                *string
	OutputFormat           *string
	LanguageCode           *string
	SyncAlignment          *bool
	ApplyTextNormalization *types.TextToDialogueRealtimeApplyTextNormalization
	Seed                   *int
	EnableLogging          *bool
}

// Realtime opens a connection to the text_to_dialogue channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *TextToDialogueConnectRequest, opts ...option.RequestOption) (*TextToDialogueSocket, error) {
	if request == nil {
		request = &TextToDialogueConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, internal.ResolveEnvironmentBaseURL(options.Environment, "Production"), c.options.BaseURL, internal.ResolveEnvironmentBaseURL(c.options.Environment, "Production"), "wss://api.elevenlabs.io")
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
		internal.EncodeURL("/v1/text-to-dialogue/stream-input"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &TextToDialogueSocket{}
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

// TextToDialogueSocket is an open connection to the text_to_dialogue channel.
//
// Stream expressive dialogue audio over a WebSocket by sending incremental text segments per registered voice.
//
// The connection uses Eleven v3 and v4 dialogue models (`model_id` must start with `eleven_v3` or `eleven_v4`). The default model is `eleven_v3_conversational`.
//
// ## Session setup
// - After connecting, the first JSON message **must** include `voices` (voice IDs to register for the session) and credentials if not already sent via headers or query string.
// - Optional `voice_settings` and `pronunciation_dictionary_locators` are only accepted on the first message.
// - For `eleven_v3_conversational` and `eleven_v4` models, only **one** voice ID may be registered. For `eleven_v3`, you may register up to **10** voices.
//
// ## Streaming text
// - Send `inputs`: an array of `{ "text", "voice_id", "new_turn"? }`. Text for the same turn is buffered until the server has enough context (at least ~40 characters and 8 words), then partial audio chunks are emitted.
// - Set `new_turn` to `true` (or switch `voice_id`) to finalize the current prosody segment and start a new speaker turn.
//
// ## Control messages
// - `flush`: force generation of any buffered text without closing the socket.
// - `close_socket`: flush remaining audio, send a final message, and close the connection.
// - `keep_alive`: reset the **20 second** receive timeout (no generation).
//
// ## Authentication
// Use the `xi-api-key` or `Authorization` header, `single_use_token` query parameter, or include `xi_api_key`, `authorization`, or `single_use_token` in the first message body (same pattern as [Text to Speech WebSocket](/docs/api-reference/text-to-speech/v-1-text-to-speech-voice-id-stream-input)). Anonymous sessions are rejected.
//
// For non-streaming dialogue over HTTP, see [Create dialogue](/docs/api-reference/text-to-dialogue/convert) and [Stream dialogue](/docs/api-reference/text-to-dialogue/stream).
type TextToDialogueSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   textToDialogueSocketHandlers
}

type textToDialogueSocketHandlers struct {
	onReceiveTextToDialogueWebsocketMessage func(*types.ReceiveTextToDialogueWebsocketMessage)
	onRaw                                   func(string)
	onUnknown                               func(json.RawMessage)
	onError                                 func(error)
	onConnected                             func()
	onDisconnected                          func(core.DisconnectReason)
}

// OnReceiveTextToDialogueWebsocketMessage registers the handler for ReceiveTextToDialogueWebsocketMessage messages, replacing any earlier one.
func (s *TextToDialogueSocket) OnReceiveTextToDialogueWebsocketMessage(handler func(*types.ReceiveTextToDialogueWebsocketMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveTextToDialogueWebsocketMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *TextToDialogueSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *TextToDialogueSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *TextToDialogueSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *TextToDialogueSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *TextToDialogueSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendTextToDialogueClientMessage sends TextToDialogueClientMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *TextToDialogueSocket) SendTextToDialogueClientMessage(ctx context.Context, message *types.TextToDialogueClientMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *TextToDialogueSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *TextToDialogueSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *TextToDialogueSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *TextToDialogueSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *TextToDialogueSocket) Close() error {
	return s.connection.Close()
}

func (s *TextToDialogueSocket) snapshot() textToDialogueSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *TextToDialogueSocket) events() internal.WebSocketEvents {
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

func (s *TextToDialogueSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *TextToDialogueSocket) receiveText(data []byte) {
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
		var message types.ReceiveTextToDialogueWebsocketMessage
		if frame.Decode(&message) {
			if handlers.onReceiveTextToDialogueWebsocketMessage != nil {
				handlers.onReceiveTextToDialogueWebsocketMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *TextToDialogueSocket) receiveBinary(_ []byte) {}
