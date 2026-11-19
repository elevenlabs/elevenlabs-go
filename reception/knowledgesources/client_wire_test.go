package knowledgesources

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
	reception "github.com/elevenlabs/elevenlabs-go/reception"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionKnowledgesources(t, 200, `[
  {
    "source_type": "file",
    "source_url": "value",
    "last_scraped_at_unix": 123,
    "original_filename": "value",
    "content_type": "value",
    "file_size_bytes": 123,
    "status": "processing",
    "error_message": "value",
    "failure_reason": "blocked_url",
    "crawl_entire_site": true,
    "scraped_pages_count": 123,
    "content_preview": "value",
    "processing_started_at_unix": 123,
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.ListKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionKnowledgesources(gotResponse, []byte(`[
  {
    "source_type": "file",
    "source_url": "value",
    "last_scraped_at_unix": 123,
    "original_filename": "value",
    "content_type": "value",
    "file_size_bytes": 123,
    "status": "processing",
    "error_message": "value",
    "failure_reason": "blocked_url",
    "crawl_entire_site": true,
    "scraped_pages_count": 123,
    "content_preview": "value",
    "processing_started_at_unix": 123,
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "source_type": "file",
    "source_url": "value",
    "last_scraped_at_unix": 123,
    "original_filename": "value",
    "content_type": "value",
    "file_size_bytes": 123,
    "status": "processing",
    "error_message": "value",
    "failure_reason": "blocked_url",
    "crawl_entire_site": true,
    "scraped_pages_count": 123,
    "content_preview": "value",
    "processing_started_at_unix": 123,
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123
  }
]`)
	}

	req := receiveWireReceptionKnowledgesources(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/knowledge-sources" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/knowledge-sources")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionKnowledgesources(t, 422, `{
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

	request := &reception.ListKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
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

	server, recorded := serveWireReceptionKnowledgesources(t, 200, `{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.GetKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionKnowledgesources(gotResponse, []byte(`{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionKnowledgesources(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/knowledge-sources/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/knowledge-sources/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionKnowledgesources(t, 422, `{
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

	request := &reception.GetKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Get(context.Background(), string("value"), request)
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

	server, recorded := serveWireReceptionKnowledgesources(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.DeleteKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionKnowledgesources(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionKnowledgesources(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/knowledge-sources/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/knowledge-sources/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionKnowledgesources(t, 422, `{
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

	request := &reception.DeleteKnowledgeSourcesRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Delete(context.Background(), string("value"), request)
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

func TestAddUrlWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionKnowledgesources(t, 200, `{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.AddUrlRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "url": "value",
  "crawl_entire_site": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.AddUrl(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionKnowledgesources(gotResponse, []byte(`{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "source_type": "file",
  "source_url": "value",
  "last_scraped_at_unix": 123,
  "original_filename": "value",
  "content_type": "value",
  "file_size_bytes": 123,
  "status": "processing",
  "error_message": "value",
  "failure_reason": "blocked_url",
  "crawl_entire_site": true,
  "scraped_pages_count": 123,
  "content_preview": "value",
  "processing_started_at_unix": 123,
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionKnowledgesources(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/knowledge-sources/urls" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/knowledge-sources/urls")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionKnowledgesources(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAddUrlWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionKnowledgesources(t, 422, `{
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

	var request reception.AddUrlRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "url": "value",
  "crawl_entire_site": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.AddUrl(context.Background(), &request)
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

type wireRequestReceptionKnowledgesources struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionKnowledgesources records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionKnowledgesources(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionKnowledgesources) {
	t.Helper()
	recorded := make(chan wireRequestReceptionKnowledgesources, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionKnowledgesources{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionKnowledgesources(t *testing.T, recorded <-chan wireRequestReceptionKnowledgesources) wireRequestReceptionKnowledgesources {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionKnowledgesources{}
	}
}

// wireJSONContainsReceptionKnowledgesources reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionKnowledgesources(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionKnowledgesources(got, want)
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

// setWireValueReceptionKnowledgesources sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionKnowledgesources(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionKnowledgesources(got, want []byte) bool {
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
