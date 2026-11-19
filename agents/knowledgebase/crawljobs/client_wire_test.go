package crawljobs

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

	knowledgebase "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseCrawljobs(t, 200, `{
  "crawl_jobs": [
    {
      "type": "discovery",
      "seed_url": "value",
      "pattern": "value",
      "exclude_patterns": [
        "value"
      ],
      "max_pages": 123,
      "auto_discover": true,
      "status": "queued",
      "pages_identified": 123,
      "pages_scraped": 123,
      "pages_skipped": 123,
      "pages_failed": 123,
      "root_folder_id": "value",
      "updated_at": 123,
      "id": "value",
      "created_at": 123
    }
  ],
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &knowledgebase.ListCrawlJobsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(gotResponse, []byte(`{
  "crawl_jobs": [
    {
      "type": "discovery",
      "seed_url": "value",
      "pattern": "value",
      "exclude_patterns": [
        "value"
      ],
      "max_pages": 123,
      "auto_discover": true,
      "status": "queued",
      "pages_identified": 123,
      "pages_scraped": 123,
      "pages_skipped": 123,
      "pages_failed": 123,
      "root_folder_id": "value",
      "updated_at": 123,
      "id": "value",
      "created_at": 123
    }
  ],
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "crawl_jobs": [
    {
      "type": "discovery",
      "seed_url": "value",
      "pattern": "value",
      "exclude_patterns": [
        "value"
      ],
      "max_pages": 123,
      "auto_discover": true,
      "status": "queued",
      "pages_identified": 123,
      "pages_scraped": 123,
      "pages_skipped": 123,
      "pages_failed": 123,
      "root_folder_id": "value",
      "updated_at": 123,
      "id": "value",
      "created_at": 123
    }
  ],
  "next_cursor": "value"
}`)
	}

	req := receiveWireAgentsKnowledgebaseCrawljobs(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/crawl" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/crawl")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseCrawljobs(t, 422, `{
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

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &knowledgebase.ListCrawlJobsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
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

	server, recorded := serveWireAgentsKnowledgebaseCrawljobs(t, 200, `{
  "id": "value",
  "type": "discovery",
  "root_folder_id": "value",
  "status": "value",
  "created_at": 123,
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "max_depth": 123,
  "max_pages": 123,
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "sitemap_urls": [
    "value"
  ],
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "auto_discover": true,
  "minimum_frequency_days": 123
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
	if !wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(gotResponse, []byte(`{
  "id": "value",
  "type": "discovery",
  "root_folder_id": "value",
  "status": "value",
  "created_at": 123,
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "type": "discovery",
  "root_folder_id": "value",
  "status": "value",
  "created_at": 123,
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)
	}

	req := receiveWireAgentsKnowledgebaseCrawljobs(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/crawl" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/crawl")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseCrawljobs(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseCrawljobs(t, 422, `{
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

	var request knowledgebase.BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "max_depth": 123,
  "max_pages": 123,
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "sitemap_urls": [
    "value"
  ],
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "auto_discover": true,
  "minimum_frequency_days": 123
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseCrawljobs(t, 200, `{
  "type": "discovery",
  "seed_url": "value",
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "max_pages": 123,
  "auto_discover": true,
  "status": "queued",
  "pages_identified": 123,
  "pages_scraped": 123,
  "pages_skipped": 123,
  "pages_failed": 123,
  "root_folder_id": "value",
  "updated_at": 123,
  "id": "value",
  "created_at": 123
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
	if !wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(gotResponse, []byte(`{
  "type": "discovery",
  "seed_url": "value",
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "max_pages": 123,
  "auto_discover": true,
  "status": "queued",
  "pages_identified": 123,
  "pages_scraped": 123,
  "pages_skipped": 123,
  "pages_failed": 123,
  "root_folder_id": "value",
  "updated_at": 123,
  "id": "value",
  "created_at": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "type": "discovery",
  "seed_url": "value",
  "pattern": "value",
  "exclude_patterns": [
    "value"
  ],
  "max_pages": 123,
  "auto_discover": true,
  "status": "queued",
  "pages_identified": 123,
  "pages_scraped": 123,
  "pages_skipped": 123,
  "pages_failed": 123,
  "root_folder_id": "value",
  "updated_at": 123,
  "id": "value",
  "created_at": 123
}`)
	}

	req := receiveWireAgentsKnowledgebaseCrawljobs(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/crawl/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/crawl/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseCrawljobs(t, 422, `{
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

func TestCancelWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseCrawljobs(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Cancel(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsKnowledgebaseCrawljobs(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/crawl/value/cancel" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/crawl/value/cancel")
	}
}

func TestCancelWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseCrawljobs(t, 422, `{
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

	_, err := client.Cancel(context.Background(), string("value"))
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

type wireRequestAgentsKnowledgebaseCrawljobs struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsKnowledgebaseCrawljobs records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsKnowledgebaseCrawljobs(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsKnowledgebaseCrawljobs) {
	t.Helper()
	recorded := make(chan wireRequestAgentsKnowledgebaseCrawljobs, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsKnowledgebaseCrawljobs{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsKnowledgebaseCrawljobs(t *testing.T, recorded <-chan wireRequestAgentsKnowledgebaseCrawljobs) wireRequestAgentsKnowledgebaseCrawljobs {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsKnowledgebaseCrawljobs{}
	}
}

// wireJSONContainsAgentsKnowledgebaseCrawljobs reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsKnowledgebaseCrawljobs(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(got, want)
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

// setWireValueAgentsKnowledgebaseCrawljobs sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsKnowledgebaseCrawljobs(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsKnowledgebaseCrawljobs(got, want []byte) bool {
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
