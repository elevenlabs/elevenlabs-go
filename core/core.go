package core

import (
	base64 "encoding/base64"
	fmt "fmt"
	http "net/http"
	url "net/url"
)

// RequestOption adapts the behavior of the client or an individual request.
type RequestOption interface {
	applyRequestOptions(*RequestOptions)
}

// RequestOptions defines all of the possible request options.
//
// This type is primarily used by the generated code and is not meant
// to be used directly; use the option package instead.
type RequestOptions struct {
	BaseURL string
	// Environment holds a value of the SDK root package's Environment type (e.g.
	// Environments.Production). It is typed interface{} here, rather than that concrete struct,
	// because core cannot import the root package without an import cycle (the root package's own
	// client/option re-export shims already import core and option). Matches the Fern oracle's own
	// convention for this field. Each endpoint reads it back through
	// internal.ResolveEnvironmentBaseURL, picking the field for that endpoint's base URL; an
	// explicit BaseURL still wins.
	Environment                interface{}
	HTTPClient                 HTTPClient
	HTTPHeader                 http.Header
	BodyProperties             map[string]interface{}
	QueryParameters            url.Values
	MaxAttempts                uint
	MaxBufSize                 int
	MaxStreamReconnectAttempts uint
	DisableStreamReconnection  bool
	DisableRetries             bool
	Username                   string
	Password                   string
	ApiKey                     *string
	WebSocketReconnect         *ReconnectOptions
	DisableWebSocketReconnect  bool
}

// NewRequestOptions returns a new *RequestOptions value.
//
// This function is primarily used by the generated code and is not meant
// to be used directly; use RequestOption instead.
func NewRequestOptions(opts ...RequestOption) *RequestOptions {
	options := &RequestOptions{
		HTTPHeader:      make(http.Header),
		BodyProperties:  make(map[string]interface{}),
		QueryParameters: make(url.Values),
	}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

// ToHeader maps the configured request options into a http.Header used
// for the request(s).
func (r *RequestOptions) ToHeader() http.Header {
	header := r.cloneHeader()
	if r.ApiKey != nil {
		header.Set("xi-api-key", fmt.Sprintf("%v", *r.ApiKey))
	}
	if header.Get("Authorization") == "" && (r.Username != "" || r.Password != "") {
		header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(r.Username+":"+r.Password)))
	}
	return header
}

func (r *RequestOptions) cloneHeader() http.Header {
	headers := r.HTTPHeader.Clone()
	if headers == nil {
		headers = make(http.Header)
	}
	return headers
}

// BaseURLOption implements the RequestOption interface.
type BaseURLOption struct {
	BaseURL string
}

func (b *BaseURLOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = b.BaseURL
}

// HTTPClientOption implements the RequestOption interface.
type HTTPClientOption struct {
	HTTPClient HTTPClient
}

func (h *HTTPClientOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPClient = h.HTTPClient
}

// HTTPHeaderOption implements the RequestOption interface.
type HTTPHeaderOption struct {
	HTTPHeader http.Header
}

func (h *HTTPHeaderOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPHeader = h.HTTPHeader
}

// BodyPropertiesOption implements the RequestOption interface.
type BodyPropertiesOption struct {
	BodyProperties map[string]interface{}
}

func (b *BodyPropertiesOption) applyRequestOptions(opts *RequestOptions) {
	opts.BodyProperties = b.BodyProperties
}

// QueryParametersOption implements the RequestOption interface.
type QueryParametersOption struct {
	QueryParameters url.Values
}

func (q *QueryParametersOption) applyRequestOptions(opts *RequestOptions) {
	opts.QueryParameters = q.QueryParameters
}

// MaxAttemptsOption implements the RequestOption interface.
type MaxAttemptsOption struct {
	MaxAttempts uint
}

func (m *MaxAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxAttempts = m.MaxAttempts
}

// MaxBufSizeOption implements the RequestOption interface.
type MaxBufSizeOption struct {
	MaxBufSize int
}

func (m *MaxBufSizeOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxBufSize = m.MaxBufSize
}

// MaxStreamReconnectAttemptsOption implements the RequestOption interface.
type MaxStreamReconnectAttemptsOption struct {
	MaxStreamReconnectAttempts uint
}

func (m *MaxStreamReconnectAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxStreamReconnectAttempts = m.MaxStreamReconnectAttempts
}

// WithoutStreamReconnectionOption implements the RequestOption interface.
type WithoutStreamReconnectionOption struct{}

func (w *WithoutStreamReconnectionOption) applyRequestOptions(opts *RequestOptions) {
	opts.DisableStreamReconnection = true
}

// WithoutRetriesOption implements the RequestOption interface.
type WithoutRetriesOption struct{}

func (w *WithoutRetriesOption) applyRequestOptions(opts *RequestOptions) {
	opts.DisableRetries = true
}

// ApiKeyOption implements the RequestOption interface.
type ApiKeyOption struct {
	ApiKey *string
}

func (a *ApiKeyOption) applyRequestOptions(opts *RequestOptions) {
	opts.ApiKey = a.ApiKey
}

// BasicAuthOption implements the RequestOption interface.
type BasicAuthOption struct {
	Username string
	Password string
}

func (b *BasicAuthOption) applyRequestOptions(opts *RequestOptions) {
	opts.Username = b.Username
	opts.Password = b.Password
}

// WebSocketReconnectOption implements the RequestOption interface.
type WebSocketReconnectOption struct {
	WebSocketReconnect ReconnectOptions
}

func (w *WebSocketReconnectOption) applyRequestOptions(opts *RequestOptions) {
	reconnect := w.WebSocketReconnect
	opts.WebSocketReconnect = &reconnect
	opts.DisableWebSocketReconnect = false
}

// WithoutWebSocketReconnectOption implements the RequestOption interface.
type WithoutWebSocketReconnectOption struct{}

func (w *WithoutWebSocketReconnectOption) applyRequestOptions(opts *RequestOptions) {
	opts.WebSocketReconnect = nil
	opts.DisableWebSocketReconnect = true
}

// EnvironmentOption implements the RequestOption interface.
type EnvironmentOption struct {
	Environment interface{}
}

func (e *EnvironmentOption) applyRequestOptions(opts *RequestOptions) {
	opts.Environment = e.Environment
}

// HTTPClient is an interface for a subset of the *http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Response is an HTTP response from an HTTP client.
type Response[T any] struct {
	StatusCode int
	Header     http.Header
	Body       T
}

// APIError is a lightweight wrapper around the standard error
// interface that preserves the status code from the RPC, if any.
type APIError struct {
	err error

	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	Body       []byte      `json:"-"`
}

// NewAPIError constructs a new API error.
func NewAPIError(statusCode int, header http.Header, body []byte, err error) *APIError {
	return &APIError{
		err:        err,
		Header:     header,
		StatusCode: statusCode,
		Body:       body,
	}
}

// Unwrap returns the underlying error. This also makes the error compatible
// with errors.As and errors.Is.
func (a *APIError) Unwrap() error {
	if a == nil {
		return nil
	}
	return a.err
}

// Error returns the API error's message.
func (a *APIError) Error() string {
	if a == nil || (a.err == nil && a.StatusCode == 0) {
		return ""
	}
	if a.err == nil {
		return fmt.Sprintf("%d", a.StatusCode)
	}
	if a.StatusCode == 0 {
		return a.err.Error()
	}
	return fmt.Sprintf("%d: %s", a.StatusCode, a.err.Error())
}
