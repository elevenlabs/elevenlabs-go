package texttospeechmulticontext

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

// TextToSpeechMultiContextConnectRequest holds the handshake parameters of the text_to_speech_multi_context channel.
type TextToSpeechMultiContextConnectRequest struct {
	VoiceId                string
	Authorization          *string
	SingleUseToken         *string
	ModelId                *string
	LanguageCode           *string
	EnableLogging          *bool
	OutputFormat           *string
	InactivityTimeout      *int
	SyncAlignment          *bool
	AutoMode               *bool
	ApplyTextNormalization *types.TextToSpeechMultiContextRealtimeApplyTextNormalization
	Seed                   *int
	EnableSsmlParsing      *bool
}

// Realtime opens a connection to the text_to_speech_multi_context channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *TextToSpeechMultiContextConnectRequest, opts ...option.RequestOption) (*TextToSpeechMultiContextSocket, error) {
	if request == nil {
		request = &TextToSpeechMultiContextConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "authorization", request.Authorization)
	internal.AddWebSocketQuery(query, "single_use_token", request.SingleUseToken)
	internal.AddWebSocketQuery(query, "model_id", request.ModelId)
	internal.AddWebSocketQuery(query, "language_code", request.LanguageCode)
	internal.AddWebSocketQuery(query, "enable_logging", request.EnableLogging)
	internal.AddWebSocketQuery(query, "output_format", request.OutputFormat)
	internal.AddWebSocketQuery(query, "inactivity_timeout", request.InactivityTimeout)
	internal.AddWebSocketQuery(query, "sync_alignment", request.SyncAlignment)
	internal.AddWebSocketQuery(query, "auto_mode", request.AutoMode)
	internal.AddWebSocketQuery(query, "apply_text_normalization", request.ApplyTextNormalization)
	internal.AddWebSocketQuery(query, "seed", request.Seed)
	internal.AddWebSocketQuery(query, "enable_ssml_parsing", request.EnableSsmlParsing)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/text-to-speech/%v/multi-stream-input", request.VoiceId),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &TextToSpeechMultiContextSocket{}
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

// TextToSpeechMultiContextSocket is an open connection to the text_to_speech_multi_context channel.
//
// The Multi-Context Text-to-Speech WebSockets API allows for generating audio from text input
// while managing multiple independent audio generation streams (contexts) over a single WebSocket connection.
// This is useful for scenarios requiring concurrent or interleaved audio generations, such as dynamic
// conversational AI applications.
//
// Each context, identified by a context id, maintains its own state. You can send text to specific
// contexts, flush them, or close them independently. A `close_socket` message can be used to terminate
// the entire connection gracefully.
//
// For more information on best practices for how to use this API, please see the [multi context websocket guide](/docs/eleven-api/guides/how-to/websockets/multi-context-web-socket).
type TextToSpeechMultiContextSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   textToSpeechMultiContextSocketHandlers
}

type textToSpeechMultiContextSocketHandlers struct {
	onReceiveMessageMulti func(*types.ReceiveMessageMulti)
	onRaw                 func(string)
	onUnknown             func(json.RawMessage)
	onError               func(error)
	onConnected           func()
	onDisconnected        func(core.DisconnectReason)
}

// OnReceiveMessageMulti registers the handler for ReceiveMessageMulti messages, replacing any earlier one.
func (s *TextToSpeechMultiContextSocket) OnReceiveMessageMulti(handler func(*types.ReceiveMessageMulti)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveMessageMulti = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *TextToSpeechMultiContextSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *TextToSpeechMultiContextSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *TextToSpeechMultiContextSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *TextToSpeechMultiContextSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *TextToSpeechMultiContextSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendMessageMulti sends SendMessageMulti as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *TextToSpeechMultiContextSocket) SendSendMessageMulti(ctx context.Context, message *types.SendMessageMulti) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *TextToSpeechMultiContextSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *TextToSpeechMultiContextSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *TextToSpeechMultiContextSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *TextToSpeechMultiContextSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *TextToSpeechMultiContextSocket) Close() error {
	return s.connection.Close()
}

func (s *TextToSpeechMultiContextSocket) snapshot() textToSpeechMultiContextSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *TextToSpeechMultiContextSocket) events() internal.WebSocketEvents {
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

func (s *TextToSpeechMultiContextSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *TextToSpeechMultiContextSocket) receiveText(data []byte) {
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
		var message types.ReceiveMessageMulti
		if frame.Decode(&message) {
			if handlers.onReceiveMessageMulti != nil {
				handlers.onReceiveMessageMulti(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *TextToSpeechMultiContextSocket) receiveBinary(_ []byte) {}
