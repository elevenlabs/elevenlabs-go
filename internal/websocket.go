package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/elevenlabs/elevenlabs-go/core"
)

// ErrWebSocketClosed is returned by a send on a WebSocket connection that was closed by the caller.
var ErrWebSocketClosed = errors.New("websocket connection is closed")

// ErrWebSocketReconnectGaveUp is returned by a send once every reconnection attempt has failed.
var ErrWebSocketReconnectGaveUp = errors.New("websocket reconnection attempts exhausted")

// ErrWebSocketConnectionLost is returned by a send after the connection dropped with reconnection off.
var ErrWebSocketConnectionLost = errors.New("websocket connection was lost")

const abnormalClosure = 1006

// WebSocketEvents receives what a WebSocket connection observes, in order, once it is started.
type WebSocketEvents struct {
	OnText         func(data []byte)
	OnBinary       func(data []byte)
	OnConnected    func()
	OnDisconnected func(reason core.DisconnectReason)
	OnError        func(err error)
}

// WebSocketConnectionParams configures a WebSocket connection. Header runs before every
// handshake, so credentials are read again on each reconnection. An error from Header stops that
// handshake: it fails the first connection, and is reported before the next attempt on a
// reconnection.
type WebSocketConnectionParams struct {
	URL        string
	Header     func() (http.Header, error)
	HTTPClient core.HTTPClient
	Reconnect  *core.ReconnectOptions
	Events     WebSocketEvents
}

type webSocketEpoch struct {
	conn *websocket.Conn
	lost chan struct{}
}

// WebSocketConnection is a WebSocket connection that reconnects after it is lost and holds sends
// made while it does.
type WebSocketConnection struct {
	params    WebSocketConnectionParams
	reconnect *core.ReconnectOptions
	events    WebSocketEvents
	queue     *webSocketEventQueue

	lifetime context.Context
	cancel   context.CancelFunc
	done     chan struct{}

	mu        sync.Mutex
	epoch     *webSocketEpoch
	ready     chan struct{}
	state     core.ReadyState
	closeErr  error
	closeOnce sync.Once
	startOnce sync.Once
}

// ConnectWebSocket dials the first connection and starts reading from it. Events are held until
// Start is called.
func ConnectWebSocket(ctx context.Context, params WebSocketConnectionParams) (*WebSocketConnection, error) {
	conn, err := dialWebSocketConnection(ctx, params)
	if err != nil {
		return nil, err
	}
	lifetime, cancel := context.WithCancel(context.Background())
	connection := &WebSocketConnection{
		params:   params,
		events:   params.Events,
		queue:    newWebSocketEventQueue(),
		lifetime: lifetime,
		cancel:   cancel,
		done:     make(chan struct{}),
		ready:    make(chan struct{}),
		state:    core.ReadyStateOpen,
	}
	if params.Reconnect != nil {
		options := params.Reconnect.WithDefaults()
		connection.reconnect = &options
	}
	epoch := connection.open(conn)
	connection.queue.push(connection.connected)
	go connection.run(epoch)
	return connection, nil
}

// Start delivers the held events and every later one to the event handlers.
func (c *WebSocketConnection) Start() {
	c.startOnce.Do(func() {
		go c.queue.drain(c.report)
	})
}

// ReadyState reports the connection's current state.
func (c *WebSocketConnection) ReadyState() core.ReadyState {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.state
}

// SendText writes a text frame, waiting for a reconnection in progress.
func (c *WebSocketConnection) SendText(ctx context.Context, data []byte) error {
	return c.write(ctx, websocket.MessageText, data)
}

// SendBinary writes a binary frame, waiting for a reconnection in progress.
func (c *WebSocketConnection) SendBinary(ctx context.Context, data []byte) error {
	return c.write(ctx, websocket.MessageBinary, data)
}

// Close closes the connection and stops reconnecting. It is safe to call more than once.
func (c *WebSocketConnection) Close() error {
	c.closeOnce.Do(func() {
		c.mu.Lock()
		epoch := c.epoch
		if c.closeErr == nil {
			c.closeErr = ErrWebSocketClosed
			c.state = core.ReadyStateClosing
		}
		c.mu.Unlock()
		if epoch != nil {
			_ = epoch.conn.Close(websocket.StatusNormalClosure, "")
		}
		c.cancel()
		<-c.done
	})
	return nil
}

func (c *WebSocketConnection) open(conn *websocket.Conn) *webSocketEpoch {
	conn.SetReadLimit(-1)
	epoch := &webSocketEpoch{conn: conn, lost: make(chan struct{})}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.epoch = epoch
	c.state = core.ReadyStateOpen
	close(c.ready)
	return epoch
}

func (c *WebSocketConnection) run(epoch *webSocketEpoch) {
	defer close(c.done)
	defer c.queue.close()
	for epoch != nil {
		reason := c.read(epoch.conn)
		reconnecting, closedByCaller := c.lose(epoch)
		if closedByCaller {
			c.finish(ErrWebSocketClosed)
			return
		}
		c.queue.push(func() { c.disconnected(reason) })
		if !reconnecting {
			c.finish(ErrWebSocketConnectionLost)
			return
		}
		epoch = c.reestablish()
	}
}

func (c *WebSocketConnection) read(conn *websocket.Conn) core.DisconnectReason {
	for {
		messageType, data, err := conn.Read(c.lifetime)
		if err != nil {
			return disconnectReason(err)
		}
		if messageType == websocket.MessageBinary {
			c.queue.push(func() { c.binary(data) })
			continue
		}
		c.queue.push(func() { c.text(data) })
	}
}

func disconnectReason(err error) core.DisconnectReason {
	var closeError websocket.CloseError
	if errors.As(err, &closeError) {
		return core.DisconnectReason{Code: int(closeError.Code), Reason: closeError.Reason}
	}
	return core.DisconnectReason{Code: abnormalClosure, Err: err}
}

func (c *WebSocketConnection) lose(epoch *webSocketEpoch) (reconnecting bool, closedByCaller bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(epoch.lost)
	c.epoch = nil
	c.ready = make(chan struct{})
	closedByCaller = c.closeErr != nil
	reconnecting = c.reconnect != nil && !closedByCaller
	if reconnecting {
		c.state = core.ReadyStateConnecting
	}
	return reconnecting, closedByCaller
}

func (c *WebSocketConnection) reestablish() *webSocketEpoch {
	for attempt := 1; attempt <= c.reconnect.MaxAttempts; attempt++ {
		if !c.wait(c.reconnect.Delay(attempt)) {
			c.finish(ErrWebSocketClosed)
			return nil
		}
		conn, err := c.dialOnce()
		if err == nil {
			epoch := c.open(conn)
			c.queue.push(c.connected)
			return epoch
		}
		if c.lifetime.Err() != nil {
			c.finish(ErrWebSocketClosed)
			return nil
		}
		c.queue.push(func() { c.report(err) })
	}
	c.finish(ErrWebSocketReconnectGaveUp)
	return nil
}

func (c *WebSocketConnection) dialOnce() (*websocket.Conn, error) {
	ctx, cancel := context.WithTimeout(c.lifetime, c.reconnect.DialTimeout)
	defer cancel()
	return dialWebSocketConnection(ctx, c.params)
}

func dialWebSocketConnection(ctx context.Context, params WebSocketConnectionParams) (*websocket.Conn, error) {
	header := http.Header{}
	if params.Header != nil {
		var err error
		header, err = params.Header()
		if err != nil {
			return nil, err
		}
	}
	return DialWebSocket(ctx, params.URL, header, params.HTTPClient)
}

func (c *WebSocketConnection) wait(delay time.Duration) bool {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-c.lifetime.Done():
		return false
	}
}

func (c *WebSocketConnection) finish(err error) {
	c.mu.Lock()
	if c.closeErr == nil {
		c.closeErr = err
	}
	c.state = core.ReadyStateClosed
	c.mu.Unlock()
	c.cancel()
}

func (c *WebSocketConnection) snapshot() (*webSocketEpoch, chan struct{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.state == core.ReadyStateClosed || c.state == core.ReadyStateClosing {
		return nil, nil, c.closeErr
	}
	return c.epoch, c.ready, nil
}

func (c *WebSocketConnection) write(ctx context.Context, messageType websocket.MessageType, data []byte) error {
	for {
		epoch, ready, err := c.snapshot()
		if err != nil {
			return err
		}
		if epoch == nil {
			if err := c.await(ctx, ready); err != nil {
				return err
			}
			continue
		}
		writeErr := epoch.conn.Write(ctx, messageType, data)
		if writeErr == nil {
			return nil
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if c.reconnect == nil {
			return writeErr
		}
		if err := c.await(ctx, epoch.lost); err != nil {
			return err
		}
	}
}

func (c *WebSocketConnection) await(ctx context.Context, signal <-chan struct{}) error {
	select {
	case <-signal:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-c.lifetime.Done():
		c.mu.Lock()
		defer c.mu.Unlock()
		return c.closeErr
	}
}

func (c *WebSocketConnection) text(data []byte) {
	if c.events.OnText != nil {
		c.events.OnText(data)
	}
}

func (c *WebSocketConnection) binary(data []byte) {
	if c.events.OnBinary != nil {
		c.events.OnBinary(data)
	}
}

func (c *WebSocketConnection) connected() {
	if c.events.OnConnected != nil {
		c.events.OnConnected()
	}
}

func (c *WebSocketConnection) disconnected(reason core.DisconnectReason) {
	if c.events.OnDisconnected != nil {
		c.events.OnDisconnected(reason)
	}
}

func (c *WebSocketConnection) report(err error) {
	if c.events.OnError == nil {
		return
	}
	defer func() {
		_ = recover()
	}()
	c.events.OnError(err)
}

type webSocketEventQueue struct {
	mu     sync.Mutex
	events []func()
	closed bool
	signal chan struct{}
}

func newWebSocketEventQueue() *webSocketEventQueue {
	return &webSocketEventQueue{signal: make(chan struct{}, 1)}
}

func (q *webSocketEventQueue) push(event func()) {
	q.mu.Lock()
	q.events = append(q.events, event)
	q.mu.Unlock()
	q.notify()
}

func (q *webSocketEventQueue) close() {
	q.mu.Lock()
	q.closed = true
	q.mu.Unlock()
	q.notify()
}

func (q *webSocketEventQueue) notify() {
	select {
	case q.signal <- struct{}{}:
	default:
	}
}

func (q *webSocketEventQueue) take() ([]func(), bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	events := q.events
	q.events = nil
	return events, q.closed
}

func (q *webSocketEventQueue) drain(report func(error)) {
	for {
		events, closed := q.take()
		for _, event := range events {
			runWebSocketEvent(event, report)
		}
		if closed && len(events) == 0 {
			return
		}
		if len(events) == 0 {
			<-q.signal
		}
	}
}

func runWebSocketEvent(event func(), report func(error)) {
	defer func() {
		if recovered := recover(); recovered != nil {
			report(fmt.Errorf("websocket handler panicked: %v", recovered))
		}
	}()
	event()
}
