package client

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

	server, recorded := serveWireDubbingClient(t, 200, `{
  "dubs": [
    {
      "dubbing_id": "value",
      "name": "value",
      "status": "value",
      "source_language": "value",
      "target_languages": [
        "value"
      ],
      "editable": true,
      "created_at": "2024-01-01T00:00:00Z",
      "media_metadata": {
        "content_type": "value",
        "duration": 9.99
      },
      "error": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(123)
	request := &types.ListDubbingRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingClient(gotResponse, []byte(`{
  "dubs": [
    {
      "dubbing_id": "value",
      "name": "value",
      "status": "value",
      "source_language": "value",
      "target_languages": [
        "value"
      ],
      "editable": true,
      "created_at": "2024-01-01T00:00:00Z",
      "media_metadata": {
        "content_type": "value",
        "duration": 9.99
      },
      "error": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "dubs": [
    {
      "dubbing_id": "value",
      "name": "value",
      "status": "value",
      "source_language": "value",
      "target_languages": [
        "value"
      ],
      "editable": true,
      "created_at": "2024-01-01T00:00:00Z",
      "media_metadata": {
        "content_type": "value",
        "duration": 9.99
      },
      "error": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireDubbingClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingClient(t, 422, `{
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

	cursorValue := string("value")
	pageSizeValue := int(123)
	request := &types.ListDubbingRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingClient(t, 200, `{
  "dubbing_id": "value",
  "name": "value",
  "status": "value",
  "source_language": "value",
  "target_languages": [
    "value"
  ],
  "editable": true,
  "created_at": "2024-01-01T00:00:00Z",
  "media_metadata": {
    "content_type": "value",
    "duration": 9.99
  },
  "error": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingClient(gotResponse, []byte(`{
  "dubbing_id": "value",
  "name": "value",
  "status": "value",
  "source_language": "value",
  "target_languages": [
    "value"
  ],
  "editable": true,
  "created_at": "2024-01-01T00:00:00Z",
  "media_metadata": {
    "content_type": "value",
    "duration": 9.99
  },
  "error": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "dubbing_id": "value",
  "name": "value",
  "status": "value",
  "source_language": "value",
  "target_languages": [
    "value"
  ],
  "editable": true,
  "created_at": "2024-01-01T00:00:00Z",
  "media_metadata": {
    "content_type": "value",
    "duration": 9.99
  },
  "error": "value"
}`)
	}

	req := receiveWireDubbingClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingClient(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"))
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

	server, recorded := serveWireDubbingClient(t, 200, `{
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
	if !wireJSONBytesEqualDubbingClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireDubbingClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/dubbing/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingClient(t, 422, `{
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

type wireRequestDubbingClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireDubbingClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireDubbingClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestDubbingClient) {
	t.Helper()
	recorded := make(chan wireRequestDubbingClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestDubbingClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireDubbingClient(t *testing.T, recorded <-chan wireRequestDubbingClient) wireRequestDubbingClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestDubbingClient{}
	}
}

// wireJSONContainsDubbingClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsDubbingClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualDubbingClient(got, want)
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

// setWireValueDubbingClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueDubbingClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualDubbingClient(got, want []byte) bool {
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
