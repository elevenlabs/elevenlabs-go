package invites

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
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
)

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceInvites(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.InviteUserRequest
	if err := json.Unmarshal([]byte(`{
  "email": "value",
  "workspace_permission": "value",
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
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
	if !wireJSONBytesEqualWorkspaceInvites(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireWorkspaceInvites(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/invites/add" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/invites/add")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceInvites(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceInvites(t, 422, `{
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

	var request workspace.InviteUserRequest
	if err := json.Unmarshal([]byte(`{
  "email": "value",
  "workspace_permission": "value",
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
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

func TestCreateBatchWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceInvites(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost
	if err := json.Unmarshal([]byte(`{
  "emails": [
    "value"
  ],
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateBatch(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceInvites(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireWorkspaceInvites(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/invites/add-bulk" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/invites/add-bulk")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceInvites(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateBatchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceInvites(t, 422, `{
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

	var request workspace.BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost
	if err := json.Unmarshal([]byte(`{
  "emails": [
    "value"
  ],
  "seat_type": "workspace_admin",
  "group_ids": [
    "value"
  ],
  "usage_limit": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateBatch(context.Background(), &request)
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

	server, recorded := serveWireWorkspaceInvites(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyDeleteExistingInvitationV1WorkspaceInvitesDelete
	if err := json.Unmarshal([]byte(`{
  "email": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Delete(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceInvites(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireWorkspaceInvites(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/workspace/invites" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/invites")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceInvites(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceInvites(t, 422, `{
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

	var request workspace.BodyDeleteExistingInvitationV1WorkspaceInvitesDelete
	if err := json.Unmarshal([]byte(`{
  "email": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Delete(context.Background(), &request)
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

type wireRequestWorkspaceInvites struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceInvites records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceInvites(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceInvites) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceInvites, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceInvites{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceInvites(t *testing.T, recorded <-chan wireRequestWorkspaceInvites) wireRequestWorkspaceInvites {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceInvites{}
	}
}

// wireJSONContainsWorkspaceInvites reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceInvites(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceInvites(got, want)
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

// setWireValueWorkspaceInvites sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceInvites(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceInvites(got, want []byte) bool {
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
