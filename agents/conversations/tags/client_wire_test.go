package tags

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

	conversations "github.com/elevenlabs/elevenlabs-go/agents/conversations"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestAssignWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsTags(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request conversations.AssignConversationTagsRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "tag_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.Assign(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/conversations/value/tags" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/tags")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsConversationsTags(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAssignWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

	var request conversations.AssignConversationTagsRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "tag_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.Assign(context.Background(), string("value"), &request)
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

func TestUnassignWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsTags(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Unassign(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/conversations/value/tags/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/tags/value2")
	}
}

func TestUnassignWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

	err := client.Unassign(context.Background(), string("value"), string("value2"))
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsTags(t, 200, `{
  "conversation_tags": [
    {
      "tag_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "title": "value",
      "description": "value",
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &conversations.ListTagsRequest{
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
	if !wireJSONBytesEqualAgentsConversationsTags(gotResponse, []byte(`{
  "conversation_tags": [
    {
      "tag_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "title": "value",
      "description": "value",
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_tags": [
    {
      "tag_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "title": "value",
      "description": "value",
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tags" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tags")
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

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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
	request := &conversations.ListTagsRequest{
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

	server, recorded := serveWireAgentsConversationsTags(t, 200, `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request conversations.CreateConversationTagRequestModel
	if err := json.Unmarshal([]byte(`{
  "title": "value",
  "description": "value"
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
	if !wireJSONBytesEqualAgentsConversationsTags(gotResponse, []byte(`{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/tags" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tags")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsConversationsTags(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

	var request conversations.CreateConversationTagRequestModel
	if err := json.Unmarshal([]byte(`{
  "title": "value",
  "description": "value"
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

	server, recorded := serveWireAgentsConversationsTags(t, 200, `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
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
	if !wireJSONBytesEqualAgentsConversationsTags(gotResponse, []byte(`{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tags/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tags/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsTags(t, 200, `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request conversations.PatchConversationTagRequestModel
	if err := json.Unmarshal([]byte(`{
  "tag_id": "value",
  "title": "value",
  "description": "value"
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
	if !wireJSONBytesEqualAgentsConversationsTags(gotResponse, []byte(`{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/tags/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tags/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsConversationsTags(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

	var request conversations.PatchConversationTagRequestModel
	if err := json.Unmarshal([]byte(`{
  "tag_id": "value",
  "title": "value",
  "description": "value"
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

	server, recorded := serveWireAgentsConversationsTags(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsConversationsTags(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/tags/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tags/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTags(t, 422, `{
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

type wireRequestAgentsConversationsTags struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsConversationsTags records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsConversationsTags(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsConversationsTags) {
	t.Helper()
	recorded := make(chan wireRequestAgentsConversationsTags, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsConversationsTags{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsConversationsTags(t *testing.T, recorded <-chan wireRequestAgentsConversationsTags) wireRequestAgentsConversationsTags {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsConversationsTags{}
	}
}

// wireJSONContainsAgentsConversationsTags reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsConversationsTags(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsConversationsTags(got, want)
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

// setWireValueAgentsConversationsTags sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsConversationsTags(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsConversationsTags(got, want []byte) bool {
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
