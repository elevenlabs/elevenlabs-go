package subtitling

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

// SubtitlingConnectRequest holds the handshake parameters of the subtitling channel.
type SubtitlingConnectRequest struct {
	TargetLanguage  string
	SourceLanguage  *string
	InputFormat     *types.SubtitlingRealtimeInputFormat
	Eagerness       *types.SubtitlingRealtimeEagerness
	ClientSessionId *string
	EnableLogging   *bool
}

// Realtime opens a connection to the subtitling channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *SubtitlingConnectRequest, opts ...option.RequestOption) (*SubtitlingSocket, error) {
	if request == nil {
		request = &SubtitlingConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "target_language", request.TargetLanguage)
	internal.AddWebSocketQuery(query, "source_language", request.SourceLanguage)
	internal.AddWebSocketQuery(query, "input_format", request.InputFormat)
	internal.AddWebSocketQuery(query, "eagerness", request.Eagerness)
	internal.AddWebSocketQuery(query, "client_session_id", request.ClientSessionId)
	internal.AddWebSocketQuery(query, "enable_logging", request.EnableLogging)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/subtitling/realtime"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &SubtitlingSocket{}
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

// SubtitlingSocket is an open connection to the subtitling channel.
//
// The Realtime Subtitling WebSocket API turns streaming speech into timed subtitles in a target language. Send source audio, and receive transcripts of the source speech and subtitle cues in the target language. Each cue arrives as `partial_subtitle` updates while it fills, then one `committed_subtitle`: to show subtitles live, render every cue message, replacing by `cue_index`; to build a subtitle file, keep only committed cues. Authenticate with the `xi-api-key` header, or from a browser, which can't set headers, by passing a bearer token as a `bearer.<token>` WebSocket subprotocol. Available to workspaces that have been granted access.
type SubtitlingSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   subtitlingSocketHandlers
}

type subtitlingSocketHandlers struct {
	onReceiveSubtitlingStreamMessage func(*types.ReceiveSubtitlingStreamMessage)
	onRaw                            func(string)
	onUnknown                        func(json.RawMessage)
	onError                          func(error)
	onConnected                      func()
	onDisconnected                   func(core.DisconnectReason)
}

// OnReceiveSubtitlingStreamMessage registers the handler for ReceiveSubtitlingStreamMessage messages, replacing any earlier one.
func (s *SubtitlingSocket) OnReceiveSubtitlingStreamMessage(handler func(*types.ReceiveSubtitlingStreamMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveSubtitlingStreamMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *SubtitlingSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *SubtitlingSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *SubtitlingSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *SubtitlingSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *SubtitlingSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendSubtitlingStreamMessage sends SendSubtitlingStreamMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *SubtitlingSocket) SendSendSubtitlingStreamMessage(ctx context.Context, message *types.SendSubtitlingStreamMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *SubtitlingSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *SubtitlingSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *SubtitlingSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *SubtitlingSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *SubtitlingSocket) Close() error {
	return s.connection.Close()
}

func (s *SubtitlingSocket) snapshot() subtitlingSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *SubtitlingSocket) events() internal.WebSocketEvents {
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

func (s *SubtitlingSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *SubtitlingSocket) receiveText(data []byte) {
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
		var message types.ReceiveSubtitlingStreamMessage
		if frame.Decode(&message) {
			if handlers.onReceiveSubtitlingStreamMessage != nil {
				handlers.onReceiveSubtitlingStreamMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *SubtitlingSocket) receiveBinary(_ []byte) {}
