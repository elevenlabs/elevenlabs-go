package texttospeech

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

// TextToSpeechConnectRequest holds the handshake parameters of the text_to_speech channel.
type TextToSpeechConnectRequest struct {
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
	ApplyTextNormalization *types.TextToSpeechRealtimeApplyTextNormalization
	Seed                   *int
	EnableSsmlParsing      *bool
}

// Realtime opens a connection to the text_to_speech channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *TextToSpeechConnectRequest, opts ...option.RequestOption) (*TextToSpeechSocket, error) {
	if request == nil {
		request = &TextToSpeechConnectRequest{}
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
		internal.EncodeURL("/v1/text-to-speech/%v/stream-input", request.VoiceId),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &TextToSpeechSocket{}
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

// TextToSpeechSocket is an open connection to the text_to_speech channel.
//
// The Text-to-Speech WebSockets API is designed to generate audio from partial text input
// while ensuring consistency throughout the generated audio. Although highly flexible,
// the WebSockets API isn't a one-size-fits-all solution. It's well-suited for scenarios where:
//   - The input text is being streamed or generated in chunks.
//   - Word-to-audio alignment information is required.
//
// However, it may not be the best choice when:
//   - The entire input text is available upfront. Given that the generations are partial,
//     some buffering is involved, which could potentially result in slightly higher latency compared
//     to a standard HTTP request.
//   - You want to quickly experiment or prototype. Working with WebSockets can be harder and more
//     complex than using a standard HTTP API, which might slow down rapid development and testing.
type TextToSpeechSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   textToSpeechSocketHandlers
}

type textToSpeechSocketHandlers struct {
	onReceiveMessage func(*types.ReceiveMessage)
	onRaw            func(string)
	onUnknown        func(json.RawMessage)
	onError          func(error)
	onConnected      func()
	onDisconnected   func(core.DisconnectReason)
}

// OnReceiveMessage registers the handler for ReceiveMessage messages, replacing any earlier one.
func (s *TextToSpeechSocket) OnReceiveMessage(handler func(*types.ReceiveMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *TextToSpeechSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *TextToSpeechSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *TextToSpeechSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *TextToSpeechSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *TextToSpeechSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendMessage sends SendMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *TextToSpeechSocket) SendSendMessage(ctx context.Context, message *types.SendMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *TextToSpeechSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *TextToSpeechSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *TextToSpeechSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *TextToSpeechSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *TextToSpeechSocket) Close() error {
	return s.connection.Close()
}

func (s *TextToSpeechSocket) snapshot() textToSpeechSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *TextToSpeechSocket) events() internal.WebSocketEvents {
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

func (s *TextToSpeechSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *TextToSpeechSocket) receiveText(data []byte) {
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
		var message types.ReceiveMessage
		if frame.Decode(&message) {
			if handlers.onReceiveMessage != nil {
				handlers.onReceiveMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *TextToSpeechSocket) receiveBinary(_ []byte) {}
