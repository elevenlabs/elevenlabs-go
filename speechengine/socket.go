package speechengine

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

// SpeechEngineConnectRequest holds the handshake parameters of the speech_engine channel.
type SpeechEngineConnectRequest struct {
	XElevenLabsSpeechEngineAuthorization string
}

// Upstream opens a connection to the speech_engine channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Upstream(ctx context.Context, request *SpeechEngineConnectRequest, opts ...option.RequestOption) (*SpeechEngineSocket, error) {
	if request == nil {
		request = &SpeechEngineConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/speech-engine/upstream"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &SpeechEngineSocket{}
	connection, err := internal.ConnectWebSocket(ctx, internal.WebSocketConnectionParams{
		URL: address,
		Header: func() (http.Header, error) {
			declared := http.Header{}
			internal.SetWebSocketHeader(declared, "X-ElevenLabs-Speech-Engine-Authorization", request.XElevenLabsSpeechEngineAuthorization)
			return c.options.HandshakeHeader(declared, options)
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

// SpeechEngineSocket is an open connection to the speech_engine channel.
//
// The Speech Engine upstream WebSocket protocol defines the interface your server must
// implement so that ElevenLabs can connect to it during a Speech Engine conversation.
// Unlike other ElevenLabs WebSocket channels where your client connects to ElevenLabs,
// the Speech Engine reverses this relationship: **ElevenLabs is the WebSocket client and
// your server is the WebSocket server**.
//
// This page shows the WebSocket API shape, however we recommend using the provided server side SDKs
// instead of implementing this yourself. The SDKs include several helper methods and automatically
// handle auth for you. You can find SDK installation instructions and
// guides in the [Speech Engine quickstart](/docs/eleven-api/guides/cookbooks/speech-engine).
//
// Configure your server's publicly reachable WebSocket URL in the `wsUrl` field when
// creating or updating a Speech Engine via the REST API. When a user starts a conversation
// with that agent, ElevenLabs will open a WebSocket connection to your server and begin
// the message exchange described below.
//
// ## Connection flow
//
// 1. A user starts a conversation with a Speech Engine agent (via the ElevenLabs client SDK or API).
// 2. ElevenLabs opens a WebSocket connection to your `wsUrl`.
// 3. ElevenLabs sends an `init` message containing the conversation ID.
// 4. As the user speaks, ElevenLabs transcribes the audio and sends `user_transcript` messages with the full conversation history.
// 5. Your server calls an LLM and streams the response back as one or more `agent_response` messages.
// 6. ElevenLabs synthesizes the text to speech and streams the audio back to the user.
// 7. Periodic `ping` messages keep the connection alive; reply with `pong`.
// 8. When the conversation ends, ElevenLabs sends a `close` message.
//
// ## Authentication
//
// Every connection from ElevenLabs includes an `X-Elevenlabs-Speech-Engine-Authorization`
// header containing a short-lived JWT. Verify this token before accepting the WebSocket
// upgrade to ensure the connection originates from ElevenLabs.
//
// The JWT is signed with **HS256** using the SHA-256 hash of your ElevenLabs API key as
// the HMAC secret, and has:
// - **Issuer** (`iss`): `https://api.elevenlabs.io/convai/speech-engine`
// - **Subject** (`sub`): `convai_speech_engine_upstream`
// - **Expiry** (`exp`): short-lived; a 60-second clock-skew leeway is applied
//
// ## Interruption handling
//
// Each `user_transcript` message carries an `event_id`. If the user speaks again before
// your server finishes responding, a new `user_transcript` arrives with a higher `event_id`.
// Cancel your in-flight LLM call and begin responding to the new transcript. Any
// `agent_response` messages sent with an outdated `event_id` are silently discarded by
// ElevenLabs.
//
// ## Streaming responses
//
// Send LLM output as a sequence of `agent_response` messages with `is_final: false` for
// each text chunk, followed by a final `agent_response` with `is_final: true` and an empty
// `content` string. ElevenLabs begins synthesizing audio as chunks arrive, minimising
// latency.
type SpeechEngineSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   speechEngineSocketHandlers
}

type speechEngineSocketHandlers struct {
	onReceiveUpstreamMessage func(*types.ReceiveUpstreamMessage)
	onRaw                    func(string)
	onUnknown                func(json.RawMessage)
	onError                  func(error)
	onConnected              func()
	onDisconnected           func(core.DisconnectReason)
}

// OnReceiveUpstreamMessage registers the handler for ReceiveUpstreamMessage messages, replacing any earlier one.
func (s *SpeechEngineSocket) OnReceiveUpstreamMessage(handler func(*types.ReceiveUpstreamMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveUpstreamMessage = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *SpeechEngineSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *SpeechEngineSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *SpeechEngineSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *SpeechEngineSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *SpeechEngineSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendSendUpstreamMessage sends SendUpstreamMessage as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *SpeechEngineSocket) SendSendUpstreamMessage(ctx context.Context, message *types.SendUpstreamMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *SpeechEngineSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *SpeechEngineSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *SpeechEngineSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *SpeechEngineSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *SpeechEngineSocket) Close() error {
	return s.connection.Close()
}

func (s *SpeechEngineSocket) snapshot() speechEngineSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *SpeechEngineSocket) events() internal.WebSocketEvents {
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

func (s *SpeechEngineSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *SpeechEngineSocket) receiveText(data []byte) {
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
		var message types.ReceiveUpstreamMessage
		if frame.Decode(&message) {
			if handlers.onReceiveUpstreamMessage != nil {
				handlers.onReceiveUpstreamMessage(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *SpeechEngineSocket) receiveBinary(_ []byte) {}
