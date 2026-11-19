package translation

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

// TranslationConnectRequest holds the handshake parameters of the translation channel.
type TranslationConnectRequest struct {
	Authorization          *string
	TargetLanguage         string
	SourceLanguage         *string
	VoiceId                string
	InputFormat            *types.TranslationRealtimeInputFormat
	OutputFormat           *types.TranslationRealtimeOutputFormat
	TtsModelId             *types.TranslationRealtimeTtsModelId
	ApplyTextNormalization *types.TranslationRealtimeApplyTextNormalization
	Eagerness              *types.TranslationRealtimeEagerness
	ClientSessionId        *string
	EnableZrm              *bool
}

// Realtime opens a connection to the translation channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *TranslationConnectRequest, opts ...option.RequestOption) (*TranslationSocket, error) {
	if request == nil {
		request = &TranslationConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "authorization", request.Authorization)
	internal.AddWebSocketQuery(query, "target_language", request.TargetLanguage)
	internal.AddWebSocketQuery(query, "source_language", request.SourceLanguage)
	internal.AddWebSocketQuery(query, "voice_id", request.VoiceId)
	internal.AddWebSocketQuery(query, "input_format", request.InputFormat)
	internal.AddWebSocketQuery(query, "output_format", request.OutputFormat)
	internal.AddWebSocketQuery(query, "tts_model_id", request.TtsModelId)
	internal.AddWebSocketQuery(query, "apply_text_normalization", request.ApplyTextNormalization)
	internal.AddWebSocketQuery(query, "eagerness", request.Eagerness)
	internal.AddWebSocketQuery(query, "client_session_id", request.ClientSessionId)
	internal.AddWebSocketQuery(query, "enable_zrm", request.EnableZrm)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/translation/realtime"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &TranslationSocket{}
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

// TranslationSocket is an open connection to the translation channel.
//
// The Realtime Translation WebSocket API streams translation audio continuously to a target language given source audio. Currently for enterprise workspaces only.
type TranslationSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   translationSocketHandlers
}

type translationSocketHandlers struct {
	onReceiveTranslateStreamMessage func(*types.ReceiveTranslateStreamMessage)
	onRaw                           func(string)
	onUnknown                       func(json.RawMessage)
	onError                         func(error)
	onConnected                     func()
	onDisconnected                  func(core.DisconnectReason)
}

// OnReceiveTranslateStreamMessage registers the handler for ReceiveTranslateStreamMessage messages, replacing any earlier one.
func (s *TranslationSocket) OnReceiveTranslateStreamMessage(handler func(*types.ReceiveTranslateStreamMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveTranslateStreamMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *TranslationSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *TranslationSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *TranslationSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *TranslationSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *TranslationSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendTranslateStreamMessage sends SendTranslateStreamMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *TranslationSocket) SendSendTranslateStreamMessage(ctx context.Context, message *types.SendTranslateStreamMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *TranslationSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *TranslationSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *TranslationSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *TranslationSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *TranslationSocket) Close() error {
	return s.connection.Close()
}

func (s *TranslationSocket) snapshot() translationSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *TranslationSocket) events() internal.WebSocketEvents {
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

func (s *TranslationSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *TranslationSocket) receiveText(data []byte) {
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
		var message types.ReceiveTranslateStreamMessage
		if frame.Decode(&message) {
			if handlers.onReceiveTranslateStreamMessage != nil {
				handlers.onReceiveTranslateStreamMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *TranslationSocket) receiveBinary(_ []byte) {}
