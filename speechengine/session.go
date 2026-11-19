// Hand-written (listed in .fernignore): the per-connection Speech Engine protocol, mirroring
// SpeechEngineSession in the Python, TypeScript, PHP, .NET and Rust SDKs.

package speechengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/coder/websocket"
)

// ErrSessionClosed is returned when a response is sent on a session whose connection has closed.
var ErrSessionClosed = errors.New("speech engine: session is closed")

// Inbound frames carry the full conversation history, so allow far more than coder/websocket's
// 32 KiB default.
const maxMessageBytes = 16 << 20

// Bound on a single outbound write, so a peer that stops reading cannot stall a response forever.
const writeTimeout = 10 * time.Second

// ConversationMessage is one message of the conversation history sent with a user transcript.
type ConversationMessage struct {
	// Role is "user" or "agent".
	Role string `json:"role"`
	// Content is the text of the message.
	Content string `json:"content"`
}

// SpeechEngineCallbacks are the handlers run for each Speech Engine session. Every field is
// optional.
//
// OnInit, OnClose, OnError and OnDisconnect run on the connection's read loop, so they should
// return promptly. OnTranscript runs on its own goroutine; OnError can therefore also be called
// from a transcript goroutine.
type SpeechEngineCallbacks struct {
	// OnInit is called once the API assigns the conversation ID.
	OnInit func(conversationID string, session *SpeechEngineSession)

	// OnTranscript is called for each user transcript, with the conversation so far. Reply with
	// session.SendResponse (or the streaming variants), passing ctx through.
	//
	// ctx is cancelled when the user interrupts (a newer transcript arrives), when the session
	// closes, and when OnTranscript returns, so pass it to the LLM call to stop work that is no
	// longer needed. A non-nil error is reported to OnError unless ctx was already cancelled.
	OnTranscript func(ctx context.Context, transcript []ConversationMessage, session *SpeechEngineSession) error

	// OnClose is called when the API sends a close message ending the conversation. The session
	// closes the connection after it returns.
	OnClose func(session *SpeechEngineSession)

	// OnDisconnect is called once when the underlying connection ends, however it ended.
	OnDisconnect func(session *SpeechEngineSession)

	// OnError is called for malformed frames, error messages from the API, and errors returned by
	// OnTranscript.
	OnError func(err error, session *SpeechEngineSession)
}

// SpeechEngineTransport is the connection a session writes to. The bundled server and
// ServeSpeechEngineConn use a github.com/coder/websocket connection; implement it to drive a
// session over another WebSocket library, then feed inbound text frames to HandleMessage and call
// HandleDisconnect when the connection ends.
type SpeechEngineTransport interface {
	// WriteText sends one text frame.
	WriteText(ctx context.Context, data []byte) error
	// Close closes the connection.
	Close() error
}

// SpeechEngineSessionOptions configures a session.
type SpeechEngineSessionOptions struct {
	// Debug logs protocol events with the "[SpeechEngine]" prefix.
	Debug bool
}

// SpeechEngineSession is one Speech Engine conversation over one WebSocket connection. Its
// methods are safe for concurrent use.
type SpeechEngineSession struct {
	transport SpeechEngineTransport
	callbacks SpeechEngineCallbacks
	debug     bool

	baseCtx    context.Context
	baseCancel context.CancelFunc
	handlers   sync.WaitGroup
	writeMu    sync.Mutex

	mu             sync.Mutex
	conversationID string
	closed         bool
	disconnected   bool
	currentEventID *int64
	cancelTurn     context.CancelFunc
}

// turn identifies the transcript a handler's context belongs to.
type turn struct {
	session *SpeechEngineSession
	eventID *int64
}

type turnKey struct{}

// NewSpeechEngineSession creates a session over a transport you manage. Most applications use
// SpeechEngineServer, SpeechEngineResource.Handler or ServeSpeechEngineConn instead.
func NewSpeechEngineSession(transport SpeechEngineTransport, callbacks SpeechEngineCallbacks, opts *SpeechEngineSessionOptions) *SpeechEngineSession {
	return newSession(context.Background(), transport, callbacks, opts)
}

// newSession creates a session whose turn contexts derive from parent, so cancelling parent
// cancels the running OnTranscript at once. Writes use their own timeout contexts.
func newSession(parent context.Context, transport SpeechEngineTransport, callbacks SpeechEngineCallbacks, opts *SpeechEngineSessionOptions) *SpeechEngineSession {
	ctx, cancel := context.WithCancel(parent)
	session := &SpeechEngineSession{
		transport:  transport,
		callbacks:  callbacks,
		baseCtx:    ctx,
		baseCancel: cancel,
	}
	if opts != nil {
		session.debug = opts.Debug
	}
	return session
}

// ServeSpeechEngineConn runs a session over an accepted github.com/coder/websocket connection
// until the connection closes or ctx is cancelled. It does not verify the request: call
// SpeechEngineResource.VerifyRequest before accepting the connection.
func ServeSpeechEngineConn(ctx context.Context, conn *websocket.Conn, callbacks SpeechEngineCallbacks, opts *SpeechEngineSessionOptions) error {
	conn.SetReadLimit(maxMessageBytes)
	// Turns derive from ctx: handler.Close, Serve shutdown or a cancelled request stops LLM work
	// immediately, not only once the read loop next returns.
	session := newSession(ctx, &coderTransport{conn: conn}, callbacks, opts)
	err := session.readLoop(ctx, conn)
	session.HandleDisconnect()
	session.Wait()
	conn.CloseNow()
	return err
}

func (s *SpeechEngineSession) readLoop(ctx context.Context, conn *websocket.Conn) error {
	for {
		typ, data, err := conn.Read(ctx)
		if err != nil {
			status := websocket.CloseStatus(err)
			if status == websocket.StatusNormalClosure || status == websocket.StatusGoingAway || ctx.Err() != nil || s.isClosed() {
				return nil
			}
			s.logf("connection ended: %v", err)
			return err
		}
		if typ == websocket.MessageText {
			s.HandleMessage(data)
		}
	}
}

// ConversationID returns the conversation ID assigned by the API, or "" before the init message.
func (s *SpeechEngineSession) ConversationID() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.conversationID
}

// IsOpen reports whether the session can still send responses.
func (s *SpeechEngineSession) IsOpen() bool {
	return !s.isClosed()
}

func (s *SpeechEngineSession) isClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}

type inboundMessage struct {
	Type           string                `json:"type"`
	ConversationID string                `json:"conversation_id"`
	UserTranscript []ConversationMessage `json:"user_transcript"`
	EventID        *int64                `json:"event_id"`
	Message        string                `json:"message"`
}

// HandleMessage processes one inbound text frame. A malformed frame is reported to OnError rather
// than ending the session. Only needed with a custom transport.
func (s *SpeechEngineSession) HandleMessage(data []byte) {
	var msg inboundMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		s.emitError(fmt.Errorf("speech engine: invalid message: %w", err))
		return
	}
	switch msg.Type {
	case "init":
		s.mu.Lock()
		s.conversationID = msg.ConversationID
		s.mu.Unlock()
		s.logf("session initialized, conversation_id=%s", msg.ConversationID)
		if s.callbacks.OnInit != nil {
			s.callbacks.OnInit(msg.ConversationID, s)
		}
	case "user_transcript":
		s.startTurn(msg.UserTranscript, msg.EventID)
	case "ping":
		if err := s.writeJSON(map[string]any{"type": "pong"}); err != nil {
			s.logf("failed to send pong: %v", err)
		}
	case "close":
		s.cancelCurrentTurn()
		if s.callbacks.OnClose != nil {
			s.callbacks.OnClose(s)
		}
		s.Close()
	case "error":
		s.emitError(errors.New(msg.Message))
	default:
		// Unknown message types are ignored so the SDK stays forward-compatible.
	}
}

// startTurn interrupts the running transcript handler, if any, and starts a new one.
func (s *SpeechEngineSession) startTurn(transcript []ConversationMessage, eventID *int64) {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	if eventID != nil && s.currentEventID != nil && *eventID == *s.currentEventID && s.cancelTurn != nil {
		s.mu.Unlock()
		s.logf("skipping duplicate transcript, event_id=%d", *eventID)
		return
	}
	if s.cancelTurn != nil {
		s.cancelTurn()
		s.logf("interrupted: cancelling the previous response for event_id=%s", formatEventID(eventID))
	}
	ctx, cancel := context.WithCancel(s.baseCtx)
	ctx = context.WithValue(ctx, turnKey{}, &turn{session: s, eventID: eventID})
	s.cancelTurn = cancel
	s.currentEventID = eventID
	s.mu.Unlock()

	s.logf("received transcript, event_id=%s, messages=%d", formatEventID(eventID), len(transcript))
	if s.callbacks.OnTranscript == nil {
		return
	}
	s.handlers.Add(1)
	go func() {
		defer s.handlers.Done()
		defer cancel()
		defer func() {
			if r := recover(); r != nil {
				s.emitError(fmt.Errorf("speech engine: OnTranscript panicked: %v", r))
			}
		}()
		if err := s.callbacks.OnTranscript(ctx, transcript, s); err != nil && ctx.Err() == nil {
			s.emitError(err)
		}
	}()
}

// SendResponse sends a complete text response for the transcript ctx belongs to, followed by the
// final frame. ctx must be the context OnTranscript received.
//
// It returns ctx's error if the user has interrupted (nothing is sent), and ErrSessionClosed once
// the session has closed. Called with any other context it logs a warning and sends nothing.
func (s *SpeechEngineSession) SendResponse(ctx context.Context, text string) error {
	if err := s.SendResponseChunk(ctx, text); err != nil {
		return err
	}
	return s.EndResponse(ctx)
}

// SendResponseChunk sends one piece of a streamed response. chunk is a string, or an LLM stream
// event whose text is extracted (see ExtractText). Chunks without text are skipped. Call
// EndResponse when the stream ends. ctx must be the context OnTranscript received.
func (s *SpeechEngineSession) SendResponseChunk(ctx context.Context, chunk any) error {
	t, err := s.checkTurn(ctx)
	if err != nil || t == nil {
		return err
	}
	text, ok := ExtractText(chunk)
	if !ok || text == "" {
		return nil
	}
	return s.sendAgentResponse(ctx, text, false, t.eventID)
}

// EndResponse sends the final frame of a streamed response.
func (s *SpeechEngineSession) EndResponse(ctx context.Context) error {
	t, err := s.checkTurn(ctx)
	if err != nil || t == nil {
		return err
	}
	return s.sendAgentResponse(ctx, "", true, t.eventID)
}

// SendResponseStream sends every chunk received from chunks (see SendResponseChunk) and then the
// final frame once the channel is closed. It stops early, returning ctx's error, if the user
// interrupts.
func (s *SpeechEngineSession) SendResponseStream(ctx context.Context, chunks <-chan any) error {
	t, err := s.checkTurn(ctx)
	if err != nil || t == nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case chunk, ok := <-chunks:
			if !ok {
				return s.EndResponse(ctx)
			}
			if err := s.SendResponseChunk(ctx, chunk); err != nil {
				return err
			}
		}
	}
}

// checkTurn returns the turn ctx belongs to, or nil (after logging) when ctx is not a transcript
// context of this session.
func (s *SpeechEngineSession) checkTurn(ctx context.Context) (*turn, error) {
	if s.isClosed() {
		return nil, ErrSessionClosed
	}
	t, _ := ctx.Value(turnKey{}).(*turn)
	if t == nil || t.session != s {
		log.Print("[SpeechEngine] SendResponse called outside an OnTranscript handler. Responses can only be sent in reply to a user transcript, using the context OnTranscript received. To have the agent speak first, set a first message in your Speech Engine conversation config.")
		return nil, nil
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return t, nil
}

type agentResponse struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	EventID *int64 `json:"event_id,omitempty"`
	IsFinal bool   `json:"is_final"`
}

func (s *SpeechEngineSession) sendAgentResponse(ctx context.Context, content string, isFinal bool, eventID *int64) error {
	// Hold the write lock across the interruption check: once a newer turn has written a frame,
	// no frame of an interrupted turn can follow it.
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.writeJSONLocked(agentResponse{Type: "agent_response", Content: content, EventID: eventID, IsFinal: isFinal})
}

func (s *SpeechEngineSession) writeJSON(value any) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	return s.writeJSONLocked(value)
}

func (s *SpeechEngineSession) writeJSONLocked(value any) error {
	if s.isClosed() {
		return ErrSessionClosed
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// The write deliberately does not use the turn's context: coder/websocket closes the whole
	// connection when a write's context is cancelled mid-frame, and an interruption must not do that.
	writeCtx, cancel := context.WithTimeout(context.Background(), writeTimeout)
	defer cancel()
	return s.transport.WriteText(writeCtx, data)
}

// Close cancels any running transcript handler and closes the connection. It is idempotent.
func (s *SpeechEngineSession) Close() {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	s.closed = true
	s.mu.Unlock()
	s.baseCancel()
	if err := s.transport.Close(); err != nil {
		s.logf("close: %v", err)
	}
}

// HandleDisconnect marks the session closed, cancels any running transcript handler and calls
// OnDisconnect once. Only needed with a custom transport.
func (s *SpeechEngineSession) HandleDisconnect() {
	s.mu.Lock()
	if s.disconnected {
		s.mu.Unlock()
		return
	}
	s.disconnected = true
	s.closed = true
	s.mu.Unlock()
	s.baseCancel()
	if s.callbacks.OnDisconnect != nil {
		s.callbacks.OnDisconnect(s)
	}
}

// Wait blocks until every running transcript handler has returned.
func (s *SpeechEngineSession) Wait() {
	s.handlers.Wait()
}

func (s *SpeechEngineSession) cancelCurrentTurn() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cancelTurn != nil {
		s.cancelTurn()
	}
}

func (s *SpeechEngineSession) emitError(err error) {
	if s.callbacks.OnError != nil {
		s.callbacks.OnError(err, s)
		return
	}
	s.logf("unhandled error: %v", err)
}

func (s *SpeechEngineSession) logf(format string, args ...any) {
	if s.debug {
		log.Printf("[SpeechEngine] "+format, args...)
	}
}

func formatEventID(eventID *int64) string {
	if eventID == nil {
		return "none"
	}
	return fmt.Sprint(*eventID)
}

// coderTransport adapts a github.com/coder/websocket connection to SpeechEngineTransport.
type coderTransport struct {
	conn *websocket.Conn
}

func (t *coderTransport) WriteText(ctx context.Context, data []byte) error {
	return t.conn.Write(ctx, websocket.MessageText, data)
}

func (t *coderTransport) Close() error {
	return t.conn.Close(websocket.StatusNormalClosure, "")
}

// ExtractText returns the text of one LLM response chunk. It accepts a string, raw JSON
// ([]byte or json.RawMessage), a decoded map, or any value that marshals to JSON (for example a
// stream event struct from an LLM SDK), and recognizes these stream shapes:
//
//   - OpenAI Responses API: {"type": "response.output_text.delta", "delta": "..."}
//   - OpenAI Chat Completions API: {"choices": [{"delta": {"content": "..."}}]}
//   - Anthropic Messages API: {"type": "content_block_delta", "delta": {"type": "text_delta", "text": "..."}}
//   - Google Gemini API: the text parts of candidates[0].content.parts, skipping thoughts
//
// ok is false when the chunk carries no text.
func ExtractText(chunk any) (text string, ok bool) {
	var event map[string]any
	switch value := chunk.(type) {
	case nil:
		return "", false
	case string:
		return value, true
	case map[string]any:
		event = value
	case json.RawMessage:
		if json.Unmarshal(value, &event) != nil {
			return "", false
		}
	case []byte:
		if json.Unmarshal(value, &event) != nil {
			return "", false
		}
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return "", false
		}
		var decoded any
		if json.Unmarshal(data, &decoded) != nil {
			return "", false
		}
		if s, isString := decoded.(string); isString {
			return s, true
		}
		event, _ = decoded.(map[string]any)
	}
	if event == nil {
		return "", false
	}

	eventType, _ := event["type"].(string)

	// OpenAI Responses API
	if eventType == "response.output_text.delta" {
		if delta, isString := event["delta"].(string); isString {
			return delta, true
		}
	}

	// OpenAI Chat Completions API
	if choices, isArray := event["choices"].([]any); isArray && len(choices) > 0 {
		if choice, isMap := choices[0].(map[string]any); isMap {
			if delta, isMap := choice["delta"].(map[string]any); isMap {
				if content, isString := delta["content"].(string); isString {
					return content, true
				}
			}
		}
	}

	// Anthropic Messages API
	if eventType == "content_block_delta" {
		if delta, isMap := event["delta"].(map[string]any); isMap && delta["type"] == "text_delta" {
			if t, isString := delta["text"].(string); isString {
				return t, true
			}
		}
	}

	// Google Gemini API
	if candidates, isArray := event["candidates"].([]any); isArray && len(candidates) > 0 {
		if candidate, isMap := candidates[0].(map[string]any); isMap {
			if content, isMap := candidate["content"].(map[string]any); isMap {
				if parts, isArray := content["parts"].([]any); isArray {
					var texts []string
					for _, p := range parts {
						part, isMap := p.(map[string]any)
						if !isMap || part["thought"] == true {
							continue
						}
						if t, isString := part["text"].(string); isString {
							texts = append(texts, t)
						}
					}
					if len(texts) > 0 {
						return strings.Join(texts, ""), true
					}
				}
			}
		}
	}
	return "", false
}
