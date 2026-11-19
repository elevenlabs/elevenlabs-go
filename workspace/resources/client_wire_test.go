package resources

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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceResources(t, 200, `{
  "resource_id": "value",
  "resource_name": "value",
  "resource_type": "voice",
  "creator_user_id": "value",
  "anonymous_access_level_override": "admin",
  "role_to_group_ids": {
    "key": [
      "value"
    ]
  },
  "share_options": [
    {
      "name": "value",
      "id": "value",
      "type": "user"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &workspace.GetResourcesRequest{}
	setWireValueWorkspaceResources(t, &request.ResourceType, "voice")
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceResources(gotResponse, []byte(`{
  "resource_id": "value",
  "resource_name": "value",
  "resource_type": "voice",
  "creator_user_id": "value",
  "anonymous_access_level_override": "admin",
  "role_to_group_ids": {
    "key": [
      "value"
    ]
  },
  "share_options": [
    {
      "name": "value",
      "id": "value",
      "type": "user"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "resource_id": "value",
  "resource_name": "value",
  "resource_type": "voice",
  "creator_user_id": "value",
  "anonymous_access_level_override": "admin",
  "role_to_group_ids": {
    "key": [
      "value"
    ]
  },
  "share_options": [
    {
      "name": "value",
      "id": "value",
      "type": "user"
    }
  ]
}`)
	}

	req := receiveWireWorkspaceResources(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/resources/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/resources/value")
	}
	if got := req.Query.Get("resource_type"); got != "voice" {
		t.Fatalf("query %s: got %q want %q", "resource_type", got, "voice")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceResources(t, 422, `{
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

	request := &workspace.GetResourcesRequest{}
	setWireValueWorkspaceResources(t, &request.ResourceType, "voice")
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

func TestShareWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceResources(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost
	if err := json.Unmarshal([]byte(`{
  "resource_id": "value",
  "role": "admin",
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Share(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceResources(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireWorkspaceResources(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/resources/value/share" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/resources/value/share")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceResources(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestShareWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceResources(t, 422, `{
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

	var request workspace.BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost
	if err := json.Unmarshal([]byte(`{
  "resource_id": "value",
  "role": "admin",
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Share(context.Background(), string("value"), &request)
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

func TestUnshareWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceResources(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost
	if err := json.Unmarshal([]byte(`{
  "resource_id": "value",
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Unshare(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceResources(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireWorkspaceResources(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/resources/value/unshare" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/resources/value/unshare")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceResources(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUnshareWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceResources(t, 422, `{
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

	var request workspace.BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost
	if err := json.Unmarshal([]byte(`{
  "resource_id": "value",
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Unshare(context.Background(), string("value"), &request)
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

type wireRequestWorkspaceResources struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceResources records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceResources(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceResources) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceResources, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceResources{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceResources(t *testing.T, recorded <-chan wireRequestWorkspaceResources) wireRequestWorkspaceResources {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceResources{}
	}
}

// wireJSONContainsWorkspaceResources reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceResources(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceResources(got, want)
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

// setWireValueWorkspaceResources sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceResources(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceResources(got, want []byte) bool {
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
