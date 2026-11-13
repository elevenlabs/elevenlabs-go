package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/coder/websocket"
	"github.com/elevenlabs/elevenlabs-go/core"
)

// WebSocketURL joins a base URL and a channel path, then merges the query parameters onto it.
func WebSocketURL(baseURL string, path string, query url.Values) (string, error) {
	parsed, err := url.Parse(strings.TrimRight(baseURL, "/") + path)
	if err != nil {
		return "", fmt.Errorf("invalid websocket url: %w", redactURLError(err))
	}
	values := parsed.Query()
	for key, items := range query {
		values[key] = items
	}
	parsed.RawQuery = values.Encode()
	return parsed.String(), nil
}

// MergeQueryValues returns the union of the given query values, where the right takes precedence.
func MergeQueryValues(sets ...url.Values) url.Values {
	merged := make(url.Values)
	for _, set := range sets {
		for key, values := range set {
			if len(values) > 0 {
				merged[key] = append([]string(nil), values...)
			}
		}
	}
	return merged
}

// AddWebSocketQuery adds a declared query parameter, skipping a nil pointer and repeating the key
// for each element of a slice.
func AddWebSocketQuery(query url.Values, key string, value interface{}) {
	for _, item := range webSocketParameterValues(value) {
		query.Add(key, item)
	}
}

// SetWebSocketQueryCredential sets a query-located credential, preferring the first
// non-empty value, so a per-call key beats the client's.
func SetWebSocketQueryCredential(query url.Values, key string, values ...string) {
	for _, value := range values {
		if value != "" {
			query.Set(key, value)
			return
		}
	}
}

// SetWebSocketHeader sets a declared header, skipping a nil pointer.
func SetWebSocketHeader(header http.Header, key string, value interface{}) {
	values := webSocketParameterValues(value)
	if len(values) > 0 {
		header.Set(key, strings.Join(values, ","))
	}
}

func webSocketParameterValues(value interface{}) []string {
	reflected := reflect.ValueOf(value)
	for reflected.Kind() == reflect.Ptr {
		if reflected.IsNil() {
			return nil
		}
		reflected = reflected.Elem()
	}
	if !reflected.IsValid() {
		return nil
	}
	if reflected.Kind() == reflect.Slice || reflected.Kind() == reflect.Array {
		values := make([]string, 0, reflected.Len())
		for index := 0; index < reflected.Len(); index++ {
			values = append(values, webSocketParameterValues(reflected.Index(index).Interface())...)
		}
		return values
	}
	return []string{fmt.Sprintf("%v", reflected.Interface())}
}

// WebSocketHTTPClient returns the first HTTP client configured on the given options.
func WebSocketHTTPClient(options ...*core.RequestOptions) core.HTTPClient {
	for _, option := range options {
		if option != nil && option.HTTPClient != nil {
			return option.HTTPClient
		}
	}
	return nil
}

// WebSocketReconnect resolves the reconnection settings, the per-call options taking precedence
// over the client's. It returns nil when reconnection is turned off.
func WebSocketReconnect(clientOptions *core.RequestOptions, callOptions *core.RequestOptions) *core.ReconnectOptions {
	for _, option := range []*core.RequestOptions{callOptions, clientOptions} {
		if option == nil {
			continue
		}
		if option.DisableWebSocketReconnect {
			return nil
		}
		if option.WebSocketReconnect != nil {
			reconnect := *option.WebSocketReconnect
			return &reconnect
		}
	}
	reconnect := core.DefaultReconnectOptions()
	return &reconnect
}

// DialWebSocket opens a WebSocket handshake. A rejected handshake is returned as *core.APIError,
// and neither error carries the URL's query or credentials.
func DialWebSocket(ctx context.Context, rawURL string, header http.Header, client core.HTTPClient) (*websocket.Conn, error) {
	options := &websocket.DialOptions{HTTPHeader: header}
	if httpClient, ok := client.(*http.Client); ok {
		options.HTTPClient = httpClient
	}
	conn, response, err := websocket.Dial(ctx, rawURL, options)
	if err == nil {
		return conn, nil
	}
	redacted := redactURLError(err)
	if response != nil && response.StatusCode != http.StatusSwitchingProtocols {
		var body []byte
		if response.Body != nil {
			body, _ = io.ReadAll(response.Body)
			response.Body.Close()
		}
		return nil, core.NewAPIError(response.StatusCode, response.Header, body, redacted)
	}
	return nil, redacted
}

func redactURLError(err error) error {
	var urlError *url.Error
	if !errors.As(err, &urlError) {
		return err
	}
	message := strings.ReplaceAll(err.Error(), urlError.URL, RedactURL(urlError.URL))
	return &redactedURLError{message: message, cause: urlError.Err}
}

type redactedURLError struct {
	message string
	cause   error
}

func (e *redactedURLError) Error() string {
	return e.message
}

func (e *redactedURLError) Unwrap() error {
	return e.cause
}

// RedactURL drops the query and user information from a URL so it can be logged.
func RedactURL(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "<invalid url>"
	}
	parsed.User = nil
	parsed.RawQuery = ""
	parsed.ForceQuery = false
	return parsed.String()
}
