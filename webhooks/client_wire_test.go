package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "webhooks": [
    {
      "name": "value",
      "webhook_id": "value",
      "webhook_url": "value",
      "is_disabled": true,
      "is_auto_disabled": true,
      "created_at_unix": 123,
      "auth_type": "hmac",
      "usage": [
        {
          "usage_type": "ConvAI Agent Settings"
        }
      ],
      "events": [
        "voice_library_removal_notice"
      ],
      "most_recent_failure_error_code": 123,
      "most_recent_failure_timestamp": 123
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	includeUsagesValue := bool(true)
	request := &types.ListWebhooksRequest{
		IncludeUsages: &includeUsagesValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "webhooks": [
    {
      "name": "value",
      "webhook_id": "value",
      "webhook_url": "value",
      "is_disabled": true,
      "is_auto_disabled": true,
      "created_at_unix": 123,
      "auth_type": "hmac",
      "usage": [
        {
          "usage_type": "ConvAI Agent Settings"
        }
      ],
      "events": [
        "voice_library_removal_notice"
      ],
      "most_recent_failure_error_code": 123,
      "most_recent_failure_timestamp": 123
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "webhooks": [
    {
      "name": "value",
      "webhook_id": "value",
      "webhook_url": "value",
      "is_disabled": true,
      "is_auto_disabled": true,
      "created_at_unix": 123,
      "auth_type": "hmac",
      "usage": [
        {
          "usage_type": "ConvAI Agent Settings"
        }
      ],
      "events": [
        "voice_library_removal_notice"
      ],
      "most_recent_failure_error_code": 123,
      "most_recent_failure_timestamp": 123
    }
  ]
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/webhooks" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks")
	}
	if got := req.Query.Get("include_usages"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_usages", got, "true")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	includeUsagesValue := bool(true)
	request := &types.ListWebhooksRequest{
		IncludeUsages: &includeUsagesValue,
	}
	_, err := client.List(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "webhook_id": "value",
  "webhook_secret": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost
	if err := json.Unmarshal([]byte(`{
  "settings": {
    "auth_type": "hmac",
    "name": "value",
    "webhook_url": "value",
    "request_headers": {
      "key": "value"
    },
    "log_redacted_fields": [
      "value"
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "webhook_id": "value",
  "webhook_secret": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "webhook_id": "value",
  "webhook_secret": "value"
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/webhooks" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWebhooks(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost
	if err := json.Unmarshal([]byte(`{
  "settings": {
    "auth_type": "hmac",
    "name": "value",
    "webhook_url": "value",
    "request_headers": {
      "key": "value"
    },
    "log_redacted_fields": [
      "value"
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch
	if err := json.Unmarshal([]byte(`{
  "webhook_id": "value",
  "is_disabled": true,
  "name": "value",
  "retry_enabled": true,
  "request_headers": {
    "key": "value"
  },
  "events": [
    "voice_library_removal_notice"
  ],
  "log_redacted_fields": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/workspace/webhooks/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWebhooks(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch
	if err := json.Unmarshal([]byte(`{
  "webhook_id": "value",
  "is_disabled": true,
  "name": "value",
  "retry_enabled": true,
  "request_headers": {
    "key": "value"
  },
  "events": [
    "voice_library_removal_notice"
  ],
  "log_redacted_fields": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/workspace/webhooks/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Delete(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestTestWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost
	if err := json.Unmarshal([]byte(`{
  "webhook_id": "value",
  "webhook_type": "project_conversion_status"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Test(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/webhooks/value/test" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks/value/test")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWebhooks(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestTestWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.BodySendTestWebhookEventV1WorkspaceWebhooksWebhookIdTestPost
	if err := json.Unmarshal([]byte(`{
  "webhook_id": "value",
  "webhook_type": "project_conversion_status"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Test(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestRedeliverWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWebhooks(t, 200, `{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Redeliver(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWebhooks(gotResponse, []byte(`{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "delivery_id": "value",
  "outcome": "delivered",
  "response_code": 123,
  "latency_ms": 123,
  "response_content_type": "value",
  "response_body_bytes": 123,
  "failure_origin": "value"
}`)
	}

	req := receiveWireWebhooks(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/webhooks/value/deliveries/value2/redeliver" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/webhooks/value/deliveries/value2/redeliver")
	}
}

func TestRedeliverWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWebhooks(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Redeliver(context.Background(), string("value"), string("value2"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

type wireRequestWebhooks struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWebhooks records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWebhooks(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWebhooks) {
	t.Helper()
	recorded := make(chan wireRequestWebhooks, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWebhooks{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if responseBody != "" {
			_, _ = w.Write([]byte(responseBody))
		}
	}))
	t.Cleanup(server.Close)
	return server, recorded
}

func receiveWireWebhooks(t *testing.T, recorded <-chan wireRequestWebhooks) wireRequestWebhooks {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWebhooks{}
	}
}

// wireJSONContainsWebhooks reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWebhooks(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWebhooks(got, want)
	}
	for key, wantValue := range wantObject {
		if key == "stream" {
			continue
		}
		if gotValue, ok := gotObject[key]; !ok || !reflect.DeepEqual(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// setWireValueWebhooks sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWebhooks(t *testing.T, target any, value any) {
	t.Helper()
	field := reflect.ValueOf(target).Elem()
	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}
	switch typed := value.(type) {
	case string:
		field.SetString(typed)
	case int:
		field.SetInt(int64(typed))
	case float64:
		field.SetFloat(typed)
	default:
		t.Fatalf("unsupported value %T", value)
	}
}

func wireJSONBytesEqualWebhooks(got, want []byte) bool {
	var gotValue any
	var wantValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		return false
	}
	if err := json.Unmarshal(want, &wantValue); err != nil {
		return false
	}
	return reflect.DeepEqual(gotValue, wantValue)
}
