package client

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

// DubbingConnectRequest holds the handshake parameters of the dubbing channel.
type DubbingConnectRequest struct {
	Authorization    *string
	TargetLanguage   string
	SourceLanguage   string
	MaxLagS          *int
	CloningStrength  *int
	InputFormat      *types.DubbingLiveInputFormat
	InputNumChannels *int
	OutputFormat     *types.DubbingLiveOutputFormat
	ClientSessionId  *string
}

// Live opens a connection to the dubbing channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Live(ctx context.Context, request *DubbingConnectRequest, opts ...option.RequestOption) (*DubbingSocket, error) {
	if request == nil {
		request = &DubbingConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, internal.ResolveEnvironmentBaseURL(options.Environment, "Production"), c.options.BaseURL, internal.ResolveEnvironmentBaseURL(c.options.Environment, "Production"), "wss://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "authorization", request.Authorization)
	internal.AddWebSocketQuery(query, "target_language", request.TargetLanguage)
	internal.AddWebSocketQuery(query, "source_language", request.SourceLanguage)
	internal.AddWebSocketQuery(query, "max_lag_s", request.MaxLagS)
	internal.AddWebSocketQuery(query, "cloning_strength", request.CloningStrength)
	internal.AddWebSocketQuery(query, "input_format", request.InputFormat)
	internal.AddWebSocketQuery(query, "input_num_channels", request.InputNumChannels)
	internal.AddWebSocketQuery(query, "output_format", request.OutputFormat)
	internal.AddWebSocketQuery(query, "client_session_id", request.ClientSessionId)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/dubbing/live"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &DubbingSocket{}
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

// DubbingSocket is an open connection to the dubbing channel.
//
// The Live Dubbing WebSocket API streams speaker-preserving dubbed audio to a target language given source audio. Production access requires the live dubbing feature on the authenticated workspace. Concurrent session limits apply; a connection rejected at the concurrency limit receives a `rate_limited` message. Input and output use the declared wire formats rather than the service's internal rendering format.
type DubbingSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   dubbingSocketHandlers
}

type dubbingSocketHandlers struct {
	onReceiveDubbingStreamMessage func(*types.ReceiveDubbingStreamMessage)
	onRaw                         func(string)
	onUnknown                     func(json.RawMessage)
	onError                       func(error)
	onConnected                   func()
	onDisconnected                func(core.DisconnectReason)
}

// OnReceiveDubbingStreamMessage registers the handler for ReceiveDubbingStreamMessage messages, replacing any earlier one.
func (s *DubbingSocket) OnReceiveDubbingStreamMessage(handler func(*types.ReceiveDubbingStreamMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveDubbingStreamMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *DubbingSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *DubbingSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *DubbingSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *DubbingSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *DubbingSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendDubbingStreamMessage sends SendDubbingStreamMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *DubbingSocket) SendSendDubbingStreamMessage(ctx context.Context, message *types.SendDubbingStreamMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *DubbingSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *DubbingSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *DubbingSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *DubbingSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *DubbingSocket) Close() error {
	return s.connection.Close()
}

func (s *DubbingSocket) snapshot() dubbingSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *DubbingSocket) events() internal.WebSocketEvents {
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

func (s *DubbingSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *DubbingSocket) receiveText(data []byte) {
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
		var message types.ReceiveDubbingStreamMessage
		if frame.Decode(&message) {
			if handlers.onReceiveDubbingStreamMessage != nil {
				handlers.onReceiveDubbingStreamMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *DubbingSocket) receiveBinary(_ []byte) {}
