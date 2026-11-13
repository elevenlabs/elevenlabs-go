package core

import (
	"math"
	"net/http"
	"time"
)

// ReadyState reports where a WebSocket connection is in its lifecycle.
type ReadyState int

const (
	// ReadyStateConnecting means the socket is down and a reconnection is in progress.
	ReadyStateConnecting ReadyState = iota
	// ReadyStateOpen means the socket is connected and frames flow both ways.
	ReadyStateOpen
	// ReadyStateClosing means Close was called and the close handshake is in progress.
	ReadyStateClosing
	// ReadyStateClosed means the socket is gone for good and every send fails.
	ReadyStateClosed
)

// String returns the name of the ready state.
func (r ReadyState) String() string {
	switch r {
	case ReadyStateConnecting:
		return "connecting"
	case ReadyStateOpen:
		return "open"
	case ReadyStateClosing:
		return "closing"
	case ReadyStateClosed:
		return "closed"
	}
	return "unknown"
}

// DisconnectReason describes why a WebSocket connection was lost.
type DisconnectReason struct {
	// Code is the WebSocket close code, or 1006 when the connection dropped without one.
	Code int
	// Reason is the close reason the server sent, if any.
	Reason string
	// Err is the transport error when the connection dropped without a close frame.
	Err error
}

// ReconnectOptions configures how a WebSocket connection is re-established after it is lost.
// A zero field takes the value from DefaultReconnectOptions.
type ReconnectOptions struct {
	// MinDelay is the wait before the first reconnection attempt.
	MinDelay time.Duration
	// MaxDelay caps the wait between attempts.
	MaxDelay time.Duration
	// GrowFactor multiplies the wait after every failed attempt.
	GrowFactor float64
	// MaxAttempts is how many attempts are made before the connection is given up.
	MaxAttempts int
	// DialTimeout bounds each reconnection attempt's handshake.
	DialTimeout time.Duration
}

// DefaultReconnectOptions returns the reconnection settings used when none are configured.
func DefaultReconnectOptions() ReconnectOptions {
	return ReconnectOptions{
		MinDelay:    time.Second,
		MaxDelay:    10 * time.Second,
		GrowFactor:  1.3,
		MaxAttempts: math.MaxInt32,
		DialTimeout: 10 * time.Second,
	}
}

// WithDefaults returns the options with every zero field replaced by its default.
func (r ReconnectOptions) WithDefaults() ReconnectOptions {
	defaults := DefaultReconnectOptions()
	if r.MinDelay <= 0 {
		r.MinDelay = defaults.MinDelay
	}
	if r.MaxDelay <= 0 {
		r.MaxDelay = defaults.MaxDelay
	}
	if r.MaxDelay < r.MinDelay {
		r.MaxDelay = r.MinDelay
	}
	if r.GrowFactor < 1 {
		r.GrowFactor = defaults.GrowFactor
	}
	if r.MaxAttempts <= 0 {
		r.MaxAttempts = defaults.MaxAttempts
	}
	if r.DialTimeout <= 0 {
		r.DialTimeout = defaults.DialTimeout
	}
	return r
}

// Delay returns the wait before the given reconnection attempt, counting from one.
func (r ReconnectOptions) Delay(attempt int) time.Duration {
	if attempt <= 1 {
		return r.MinDelay
	}
	delay := float64(r.MinDelay) * math.Pow(r.GrowFactor, float64(attempt-1))
	if delay >= float64(r.MaxDelay) {
		return r.MaxDelay
	}
	return time.Duration(delay)
}

// HandshakeHeader builds the headers of a WebSocket handshake. The connect call's headers rank
// highest, the channel's declared headers among them, and a header sent there with no value is
// removed. The client's credentials come next, then the client's own headers. Header names are
// compared case-insensitively.
func (r *RequestOptions) HandshakeHeader(declared http.Header, call *RequestOptions) (http.Header, error) {
	header := canonicalHandshakeHeader(r.HTTPHeader)
	fallback := header.Get("Authorization")
	header.Del("Authorization")
	for key, values := range r.handshakeCredentials() {
		header[key] = values
	}
	perConnect := canonicalHandshakeHeader(declared)
	for key, values := range call.handshakeCallHeader() {
		perConnect[key] = values
	}
	for key, values := range perConnect {
		if hasHandshakeHeaderValue(values) {
			header[key] = values
		} else {
			header.Del(key)
		}
	}
	if _, named := perConnect["Authorization"]; named || header.Get("Authorization") != "" {
		return header, nil
	}
	if fallback != "" {
		header.Set("Authorization", fallback)
	}
	return header, nil
}

func (r *RequestOptions) handshakeCredentials() http.Header {
	options := *r
	options.HTTPHeader = nil
	return options.ToHeader()
}

func (r *RequestOptions) handshakeCallHeader() http.Header {
	options := *r
	options.HTTPHeader = canonicalHandshakeHeader(r.HTTPHeader)
	return options.ToHeader()
}

func canonicalHandshakeHeader(source http.Header) http.Header {
	header := make(http.Header, len(source))
	for key, values := range source {
		canonical := http.CanonicalHeaderKey(key)
		header[canonical] = append(header[canonical], values...)
	}
	return header
}

func hasHandshakeHeaderValue(values []string) bool {
	for _, value := range values {
		if value != "" {
			return true
		}
	}
	return false
}
