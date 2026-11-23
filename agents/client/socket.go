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

// AgentsConnectRequest holds the handshake parameters of the agents channel.
type AgentsConnectRequest struct {
	AgentId string
}

// Conversation opens a connection to the agents channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Conversation(ctx context.Context, request *AgentsConnectRequest, opts ...option.RequestOption) (*AgentsSocket, error) {
	if request == nil {
		request = &AgentsConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, internal.ResolveEnvironmentBaseURL(options.Environment, "Production"), c.options.BaseURL, internal.ResolveEnvironmentBaseURL(c.options.Environment, "Production"), "wss://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "agent_id", request.AgentId)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/convai/conversation"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &AgentsSocket{}
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

// AgentsSocket is an open connection to the agents channel.
//
// Establish a WebSocket connection for real-time conversations with an AI agent.
type AgentsSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   agentsSocketHandlers
}

type agentsSocketHandlers struct {
	onSubscribeEvent func(*types.SubscribeEvent)
	onRaw            func(string)
	onUnknown        func(json.RawMessage)
	onError          func(error)
	onConnected      func()
	onDisconnected   func(core.DisconnectReason)
}

// OnSubscribeEvent registers the handler for SubscribeEvent messages, replacing any earlier one.
func (s *AgentsSocket) OnSubscribeEvent(handler func(*types.SubscribeEvent)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onSubscribeEvent = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *AgentsSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *AgentsSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *AgentsSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *AgentsSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *AgentsSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendPublishEvent sends PublishEvent as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *AgentsSocket) SendPublishEvent(ctx context.Context, message *types.PublishEvent) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *AgentsSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *AgentsSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *AgentsSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *AgentsSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *AgentsSocket) Close() error {
	return s.connection.Close()
}

func (s *AgentsSocket) snapshot() agentsSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *AgentsSocket) events() internal.WebSocketEvents {
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

func (s *AgentsSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *AgentsSocket) receiveText(data []byte) {
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
		var message types.SubscribeEvent
		if frame.Decode(&message) {
			if handlers.onSubscribeEvent != nil {
				handlers.onSubscribeEvent(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *AgentsSocket) receiveBinary(_ []byte) {}
