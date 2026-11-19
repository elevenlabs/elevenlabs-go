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
	dubbing "github.com/elevenlabs/elevenlabs-go/dubbing"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectClient(t, 200, `{
  "projects": [
    {
      "project_id": "value",
      "status": "queued",
      "reference": "value",
      "source_language": "value",
      "model_id": "value",
      "media": {
        "filename": "value",
        "duration_s": 9.99,
        "has_video": true,
        "mime_type": "value"
      },
      "language_ids": [
        "value"
      ],
      "webhook_ids": [
        "value"
      ],
      "revision": 123,
      "error": {
        "code": "value",
        "message": "value",
        "retryable": true
      },
      "warnings": [
        {
          "type": "voices_not_permitted",
          "speaker_ids": [
            "value"
          ],
          "message": "value"
        }
      ],
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(123)
	statusValue := string("value")
	request := &dubbing.ListProjectRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
		Status:   &statusValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectClient(gotResponse, []byte(`{
  "projects": [
    {
      "project_id": "value",
      "status": "queued",
      "reference": "value",
      "source_language": "value",
      "model_id": "value",
      "media": {
        "filename": "value",
        "duration_s": 9.99,
        "has_video": true,
        "mime_type": "value"
      },
      "language_ids": [
        "value"
      ],
      "webhook_ids": [
        "value"
      ],
      "revision": 123,
      "error": {
        "code": "value",
        "message": "value",
        "retryable": true
      },
      "warnings": [
        {
          "type": "voices_not_permitted",
          "speaker_ids": [
            "value"
          ],
          "message": "value"
        }
      ],
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "projects": [
    {
      "project_id": "value",
      "status": "queued",
      "reference": "value",
      "source_language": "value",
      "model_id": "value",
      "media": {
        "filename": "value",
        "duration_s": 9.99,
        "has_video": true,
        "mime_type": "value"
      },
      "language_ids": [
        "value"
      ],
      "webhook_ids": [
        "value"
      ],
      "revision": 123,
      "error": {
        "code": "value",
        "message": "value",
        "retryable": true
      },
      "warnings": [
        {
          "type": "voices_not_permitted",
          "speaker_ids": [
            "value"
          ],
          "message": "value"
        }
      ],
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "next_cursor": "value"
}`)
	}

	req := receiveWireDubbingProjectClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("status"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "status", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectClient(t, 422, `{
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
	statusValue := string("value")
	request := &dubbing.ListProjectRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
		Status:   &statusValue,
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

	server, recorded := serveWireDubbingProjectClient(t, 200, `{
  "project_id": "value",
  "status": "queued",
  "reference": "value",
  "source_language": "value",
  "model_id": "value",
  "media": {
    "filename": "value",
    "duration_s": 9.99,
    "has_video": true,
    "mime_type": "value"
  },
  "language_ids": [
    "value"
  ],
  "webhook_ids": [
    "value"
  ],
  "revision": 123,
  "error": {
    "code": "value",
    "message": "value",
    "retryable": true
  },
  "warnings": [
    {
      "type": "voices_not_permitted",
      "speaker_ids": [
        "value"
      ],
      "message": "value"
    }
  ],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
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
	if !wireJSONBytesEqualDubbingProjectClient(gotResponse, []byte(`{
  "project_id": "value",
  "status": "queued",
  "reference": "value",
  "source_language": "value",
  "model_id": "value",
  "media": {
    "filename": "value",
    "duration_s": 9.99,
    "has_video": true,
    "mime_type": "value"
  },
  "language_ids": [
    "value"
  ],
  "webhook_ids": [
    "value"
  ],
  "revision": 123,
  "error": {
    "code": "value",
    "message": "value",
    "retryable": true
  },
  "warnings": [
    {
      "type": "voices_not_permitted",
      "speaker_ids": [
        "value"
      ],
      "message": "value"
    }
  ],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "project_id": "value",
  "status": "queued",
  "reference": "value",
  "source_language": "value",
  "model_id": "value",
  "media": {
    "filename": "value",
    "duration_s": 9.99,
    "has_video": true,
    "mime_type": "value"
  },
  "language_ids": [
    "value"
  ],
  "webhook_ids": [
    "value"
  ],
  "revision": 123,
  "error": {
    "code": "value",
    "message": "value",
    "retryable": true
  },
  "warnings": [
    {
      "type": "voices_not_permitted",
      "speaker_ids": [
        "value"
      ],
      "message": "value"
    }
  ],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}`)
	}

	req := receiveWireDubbingProjectClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectClient(t, 422, `{
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

	server, recorded := serveWireDubbingProjectClient(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireDubbingProjectClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/dubbing/project/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectClient(t, 422, `{
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

	err := client.Delete(context.Background(), string("value"))
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

type wireRequestDubbingProjectClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireDubbingProjectClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireDubbingProjectClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestDubbingProjectClient) {
	t.Helper()
	recorded := make(chan wireRequestDubbingProjectClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestDubbingProjectClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireDubbingProjectClient(t *testing.T, recorded <-chan wireRequestDubbingProjectClient) wireRequestDubbingProjectClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestDubbingProjectClient{}
	}
}

// wireJSONContainsDubbingProjectClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsDubbingProjectClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualDubbingProjectClient(got, want)
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

// setWireValueDubbingProjectClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueDubbingProjectClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualDubbingProjectClient(got, want []byte) bool {
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
