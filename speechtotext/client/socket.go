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

// SpeechToTextConnectRequest holds the handshake parameters of the speech_to_text channel.
type SpeechToTextConnectRequest struct {
	ModelId                  types.SpeechToTextRealtimeModelId
	Token                    *string
	AudioFormat              *string
	LanguageCode             *string
	SecondaryLanguages       []string
	CommitStrategy           *types.SpeechToTextRealtimeCommitStrategy
	VadThreshold             *float64
	VadSilenceThresholdSecs  *float64
	MinSpeechDurationMs      *int
	MinSilenceDurationMs     *int
	IncludeTimestamps        *bool
	IncludeLanguageDetection *bool
	Keyterms                 []string
	NoVerbatim               *bool
	MultilangMode            *bool
	Diarize                  *bool
	InferSpeakerNames        *bool
	EntityDetection          *types.V1SpeechToTextRealtimeSpeechToTextRequestEntityDetection
	EntityRedaction          *types.V1SpeechToTextRealtimeSpeechToTextRequestEntityRedaction
	EntityRedactionMode      *types.SpeechToTextRealtimeEntityRedactionMode
	TranscriptEdit           *string
	FilterBackgroundAudio    *bool
	KeepaliveIntervalMs      *int
	EnableLogging            *bool
}

// Realtime opens a connection to the speech_to_text channel. The client's
// credentials and the request's parameters are sent on the handshake, and again on every
// reconnection. Register handlers on the returned socket, then call Start.
func (c *Client) Realtime(ctx context.Context, request *SpeechToTextConnectRequest, opts ...option.RequestOption) (*SpeechToTextSocket, error) {
	if request == nil {
		request = &SpeechToTextConnectRequest{}
	}
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io")
	query := url.Values{}
	internal.AddWebSocketQuery(query, "model_id", request.ModelId)
	internal.AddWebSocketQuery(query, "token", request.Token)
	internal.AddWebSocketQuery(query, "audio_format", request.AudioFormat)
	internal.AddWebSocketQuery(query, "language_code", request.LanguageCode)
	internal.AddWebSocketQuery(query, "secondary_languages", request.SecondaryLanguages)
	internal.AddWebSocketQuery(query, "commit_strategy", request.CommitStrategy)
	internal.AddWebSocketQuery(query, "vad_threshold", request.VadThreshold)
	internal.AddWebSocketQuery(query, "vad_silence_threshold_secs", request.VadSilenceThresholdSecs)
	internal.AddWebSocketQuery(query, "min_speech_duration_ms", request.MinSpeechDurationMs)
	internal.AddWebSocketQuery(query, "min_silence_duration_ms", request.MinSilenceDurationMs)
	internal.AddWebSocketQuery(query, "include_timestamps", request.IncludeTimestamps)
	internal.AddWebSocketQuery(query, "include_language_detection", request.IncludeLanguageDetection)
	internal.AddWebSocketQuery(query, "keyterms", request.Keyterms)
	internal.AddWebSocketQuery(query, "no_verbatim", request.NoVerbatim)
	internal.AddWebSocketQuery(query, "multilang_mode", request.MultilangMode)
	internal.AddWebSocketQuery(query, "diarize", request.Diarize)
	internal.AddWebSocketQuery(query, "infer_speaker_names", request.InferSpeakerNames)
	internal.AddWebSocketQuery(query, "entity_detection", request.EntityDetection)
	internal.AddWebSocketQuery(query, "entity_redaction", request.EntityRedaction)
	internal.AddWebSocketQuery(query, "entity_redaction_mode", request.EntityRedactionMode)
	internal.AddWebSocketQuery(query, "transcript_edit", request.TranscriptEdit)
	internal.AddWebSocketQuery(query, "filter_background_audio", request.FilterBackgroundAudio)
	internal.AddWebSocketQuery(query, "keepalive_interval_ms", request.KeepaliveIntervalMs)
	internal.AddWebSocketQuery(query, "enable_logging", request.EnableLogging)
	address, err := internal.WebSocketURL(
		baseURL,
		internal.EncodeURL("/v1/speech-to-text/realtime"),
		internal.MergeQueryValues(c.options.QueryParameters, query, options.QueryParameters),
	)
	if err != nil {
		return nil, err
	}
	socket := &SpeechToTextSocket{}
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

// SpeechToTextSocket is an open connection to the speech_to_text channel.
//
// Realtime speech-to-text transcription service. This WebSocket API enables streaming audio input and receiving transcription results.
//
// ## Event Flow
// - Audio chunks are sent as `input_audio_chunk` messages
// - Transcription results are streamed back as `partial_transcript` (interim) and `committed_transcript` (stable/final for that segment)
// - Supports manual commit (the default), VAD-based automatic commit, or turn prediction for models that predict turns
//
// Authentication is done either by providing a valid API key in the `xi-api-key` header or by providing a valid token in the `token` query parameter. Tokens can be generated from the [single use token endpoint](/docs/api-reference/tokens/create). Use tokens if you want to transcribe audio from the client side.
type SpeechToTextSocket struct {
	connection *internal.WebSocketConnection
	mu         sync.RWMutex
	handlers   speechToTextSocketHandlers
}

type speechToTextSocketHandlers struct {
	onReceiveTranscription func(*types.ReceiveTranscription)
	onRaw                  func(string)
	onUnknown              func(json.RawMessage)
	onError                func(error)
	onConnected            func()
	onDisconnected         func(core.DisconnectReason)
}

// OnReceiveTranscription registers the handler for ReceiveTranscription messages, replacing any earlier one.
func (s *SpeechToTextSocket) OnReceiveTranscription(handler func(*types.ReceiveTranscription)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onReceiveTranscription = handler
}

// OnRaw registers a handler that receives every text frame verbatim, before it is dispatched.
func (s *SpeechToTextSocket) OnRaw(handler func(string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onRaw = handler
}

// OnUnknown registers a handler for JSON frames that no message handler accepts.
func (s *SpeechToTextSocket) OnUnknown(handler func(json.RawMessage)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onUnknown = handler
}

// OnError registers a handler for frames that are not valid JSON, handlers that panic and failed
// reconnection attempts. The connection keeps reading after each one.
func (s *SpeechToTextSocket) OnError(handler func(error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onError = handler
}

// OnConnected registers a handler called when Start begins delivery and after every reconnection.
func (s *SpeechToTextSocket) OnConnected(handler func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onConnected = handler
}

// OnDisconnected registers a handler called whenever the connection is lost.
func (s *SpeechToTextSocket) OnDisconnected(handler func(core.DisconnectReason)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers.onDisconnected = handler
}

// SendInputAudioChunk sends InputAudioChunk as a JSON text frame. While the connection is
// reconnecting, it waits for the connection to come back.
func (s *SpeechToTextSocket) SendInputAudioChunk(ctx context.Context, message *types.InputAudioChunk) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.connection.SendText(ctx, data)
}

// SendText sends a text frame verbatim.
func (s *SpeechToTextSocket) SendText(ctx context.Context, text string) error {
	return s.connection.SendText(ctx, []byte(text))
}

// SendBinary sends a binary frame verbatim.
func (s *SpeechToTextSocket) SendBinary(ctx context.Context, data []byte) error {
	return s.connection.SendBinary(ctx, data)
}

// Start delivers received frames and connection events to the registered handlers, one at a time
// and in arrival order. Frames that arrive before Start are held, not dropped.
func (s *SpeechToTextSocket) Start() {
	s.connection.Start()
}

// ReadyState reports whether the connection is open, reconnecting or closed.
func (s *SpeechToTextSocket) ReadyState() core.ReadyState {
	return s.connection.ReadyState()
}

// Close closes the connection and stops reconnecting. Sends waiting on a reconnection fail.
func (s *SpeechToTextSocket) Close() error {
	return s.connection.Close()
}

func (s *SpeechToTextSocket) snapshot() speechToTextSocketHandlers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.handlers
}

func (s *SpeechToTextSocket) events() internal.WebSocketEvents {
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

func (s *SpeechToTextSocket) reportError(err error) {
	if handler := s.snapshot().onError; handler != nil {
		handler(err)
	}
}

func (s *SpeechToTextSocket) receiveText(data []byte) {
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
		var message types.ReceiveTranscription
		if frame.Decode(&message) {
			if handlers.onReceiveTranscription != nil {
				handlers.onReceiveTranscription(&message)
			}
			return
		}
	}
	if handlers.onUnknown != nil {
		handlers.onUnknown(frame.Raw)
	}
}

func (s *SpeechToTextSocket) receiveBinary(_ []byte) {}
