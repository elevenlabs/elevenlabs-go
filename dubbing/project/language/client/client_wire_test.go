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
	project "github.com/elevenlabs/elevenlabs-go/dubbing/project"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectLanguageClient(t, 200, `{
  "languages": [
    {
      "language_id": "value",
      "project_id": "value",
      "target_language": "value",
      "status": "queued",
      "model_id": "value",
      "voice_settings": {
        "cloning_strength": 10
      },
      "outputs": {
        "lossless_audio": "value",
        "subtitles": "value"
      },
      "revision": 123,
      "output_revision": 123,
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
	request := &project.ListLanguageRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
		Status:   &statusValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageClient(gotResponse, []byte(`{
  "languages": [
    {
      "language_id": "value",
      "project_id": "value",
      "target_language": "value",
      "status": "queued",
      "model_id": "value",
      "voice_settings": {
        "cloning_strength": 10
      },
      "outputs": {
        "lossless_audio": "value",
        "subtitles": "value"
      },
      "revision": 123,
      "output_revision": 123,
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
  "languages": [
    {
      "language_id": "value",
      "project_id": "value",
      "target_language": "value",
      "status": "queued",
      "model_id": "value",
      "voice_settings": {
        "cloning_strength": 10
      },
      "outputs": {
        "lossless_audio": "value",
        "subtitles": "value"
      },
      "revision": 123,
      "output_revision": 123,
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

	req := receiveWireDubbingProjectLanguageClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project/value/language" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language")
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

	server, _ := serveWireDubbingProjectLanguageClient(t, 422, `{
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
	request := &project.ListLanguageRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
		Status:   &statusValue,
	}
	_, err := client.List(context.Background(), string("value"), request)
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

	server, recorded := serveWireDubbingProjectLanguageClient(t, 201, `{
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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

	var request project.BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "target_language": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "translations": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageClient(gotResponse, []byte(`{
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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

	req := receiveWireDubbingProjectLanguageClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/dubbing/project/value/language" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectLanguageClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageClient(t, 422, `{
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

	var request project.BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "target_language": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "translations": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), string("value"), &request)
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

	server, recorded := serveWireDubbingProjectLanguageClient(t, 200, `{
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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

	result, err := client.Get(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageClient(gotResponse, []byte(`{
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
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

	req := receiveWireDubbingProjectLanguageClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageClient(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"), string("value2"))
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

	server, recorded := serveWireDubbingProjectLanguageClient(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireDubbingProjectLanguageClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageClient(t, 422, `{
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

	err := client.Delete(context.Background(), string("value"), string("value2"))
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

type wireRequestDubbingProjectLanguageClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireDubbingProjectLanguageClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireDubbingProjectLanguageClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestDubbingProjectLanguageClient) {
	t.Helper()
	recorded := make(chan wireRequestDubbingProjectLanguageClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestDubbingProjectLanguageClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireDubbingProjectLanguageClient(t *testing.T, recorded <-chan wireRequestDubbingProjectLanguageClient) wireRequestDubbingProjectLanguageClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestDubbingProjectLanguageClient{}
	}
}

// wireJSONContainsDubbingProjectLanguageClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsDubbingProjectLanguageClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualDubbingProjectLanguageClient(got, want)
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

// setWireValueDubbingProjectLanguageClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueDubbingProjectLanguageClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualDubbingProjectLanguageClient(got, want []byte) bool {
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
