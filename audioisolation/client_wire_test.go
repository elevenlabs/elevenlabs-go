package audioisolation

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

	server, recorded := serveWireAudioisolation(t, 200, `{
  "items": [
    {
      "id": "value",
      "title": "value",
      "created_at_unix": 123,
      "format": "value",
      "duration_seconds": 9.99,
      "download_url": "value",
      "icon_url": "value",
      "source_video_url": "value",
      "supports_video": true,
      "processing": true,
      "video_processing_failed": true,
      "preview_b64": "value"
    }
  ],
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(123)
	pageValue := int(123)
	searchValue := string("value")
	request := &types.ListAudioIsolationRequest{
		PageSize: &pageSizeValue,
		Page:     &pageValue,
		Search:   &searchValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAudioisolation(gotResponse, []byte(`{
  "items": [
    {
      "id": "value",
      "title": "value",
      "created_at_unix": 123,
      "format": "value",
      "duration_seconds": 9.99,
      "download_url": "value",
      "icon_url": "value",
      "source_video_url": "value",
      "supports_video": true,
      "processing": true,
      "video_processing_failed": true,
      "preview_b64": "value"
    }
  ],
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "items": [
    {
      "id": "value",
      "title": "value",
      "created_at_unix": 123,
      "format": "value",
      "duration_seconds": 9.99,
      "download_url": "value",
      "icon_url": "value",
      "source_video_url": "value",
      "supports_video": true,
      "processing": true,
      "video_processing_failed": true,
      "preview_b64": "value"
    }
  ],
  "has_more": true
}`)
	}

	req := receiveWireAudioisolation(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/audio-isolation/history" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/audio-isolation/history")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("page"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAudioisolation(t, 422, `{
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

	pageSizeValue := int(123)
	pageValue := int(123)
	searchValue := string("value")
	request := &types.ListAudioIsolationRequest{
		PageSize: &pageSizeValue,
		Page:     &pageValue,
		Search:   &searchValue,
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAudioisolation(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAudioisolation(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/audio-isolation/history/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/audio-isolation/history/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAudioisolation(t, 422, `{
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

type wireRequestAudioisolation struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAudioisolation records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAudioisolation(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAudioisolation) {
	t.Helper()
	recorded := make(chan wireRequestAudioisolation, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAudioisolation{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAudioisolation(t *testing.T, recorded <-chan wireRequestAudioisolation) wireRequestAudioisolation {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAudioisolation{}
	}
}

// wireJSONContainsAudioisolation reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAudioisolation(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAudioisolation(got, want)
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

// setWireValueAudioisolation sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAudioisolation(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAudioisolation(got, want []byte) bool {
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
