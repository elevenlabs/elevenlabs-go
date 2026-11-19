package attention

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

	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsAttention(t, 200, `{
  "attention": [
    {
      "attention_id": "value",
      "campaign_id": "value",
      "reason": "value",
      "detail": "value",
      "status": "open",
      "retryable": true,
      "resolution": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	campaignIdValue := string("value")
	includeResolvedValue := bool(true)
	pageSizeValue := int(123)
	cursorValue := string("value")
	request := &outboundcampaigns.ListAttentionRequest{
		CampaignId:      &campaignIdValue,
		IncludeResolved: &includeResolvedValue,
		PageSize:        &pageSizeValue,
		Cursor:          &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsAttention(gotResponse, []byte(`{
  "attention": [
    {
      "attention_id": "value",
      "campaign_id": "value",
      "reason": "value",
      "detail": "value",
      "status": "open",
      "retryable": true,
      "resolution": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "attention": [
    {
      "attention_id": "value",
      "campaign_id": "value",
      "reason": "value",
      "detail": "value",
      "status": "open",
      "retryable": true,
      "resolution": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsOutboundcampaignsAttention(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/attention" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/attention")
	}
	if got := req.Query.Get("campaign_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "campaign_id", got, "value")
	}
	if got := req.Query.Get("include_resolved"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_resolved", got, "true")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsAttention(t, 422, `{
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

	campaignIdValue := string("value")
	includeResolvedValue := bool(true)
	pageSizeValue := int(123)
	cursorValue := string("value")
	request := &outboundcampaigns.ListAttentionRequest{
		CampaignId:      &campaignIdValue,
		IncludeResolved: &includeResolvedValue,
		PageSize:        &pageSizeValue,
		Cursor:          &cursorValue,
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

func TestRetryWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsAttention(t, 200, `{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Retry(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsAttention(gotResponse, []byte(`{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireAgentsOutboundcampaignsAttention(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/outbound-campaigns/attention/value/retry" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/attention/value/retry")
	}
}

func TestRetryWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsAttention(t, 422, `{
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

	_, err := client.Retry(context.Background(), string("value"))
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

func TestResolveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsAttention(t, 200, `{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request outboundcampaigns.AttentionResolutionRequest
	if err := json.Unmarshal([]byte(`{
  "attention_id": "value",
  "resolution": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Resolve(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsAttention(gotResponse, []byte(`{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "attention_id": "value",
  "campaign_id": "value",
  "reason": "value",
  "detail": "value",
  "status": "open",
  "retryable": true,
  "resolution": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireAgentsOutboundcampaignsAttention(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/outbound-campaigns/attention/value/resolve" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/attention/value/resolve")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsOutboundcampaignsAttention(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestResolveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsAttention(t, 422, `{
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

	var request outboundcampaigns.AttentionResolutionRequest
	if err := json.Unmarshal([]byte(`{
  "attention_id": "value",
  "resolution": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Resolve(context.Background(), string("value"), &request)
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

type wireRequestAgentsOutboundcampaignsAttention struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsOutboundcampaignsAttention records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsOutboundcampaignsAttention(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsOutboundcampaignsAttention) {
	t.Helper()
	recorded := make(chan wireRequestAgentsOutboundcampaignsAttention, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsOutboundcampaignsAttention{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsOutboundcampaignsAttention(t *testing.T, recorded <-chan wireRequestAgentsOutboundcampaignsAttention) wireRequestAgentsOutboundcampaignsAttention {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsOutboundcampaignsAttention{}
	}
}

// wireJSONContainsAgentsOutboundcampaignsAttention reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsOutboundcampaignsAttention(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsOutboundcampaignsAttention(got, want)
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

// setWireValueAgentsOutboundcampaignsAttention sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsOutboundcampaignsAttention(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsOutboundcampaignsAttention(got, want []byte) bool {
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
