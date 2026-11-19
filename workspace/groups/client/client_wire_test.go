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
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceGroupsClient(t, 200, `{
  "key": {
    "name": "value",
    "id": "value",
    "members": [
      "value"
    ],
    "permissions": [
      "text_to_speech"
    ],
    "group_usage_limit": 123,
    "group_pvc_limit": 123,
    "group_concurrency_limit": 123,
    "group_dubbing_concurrency_limit": 123,
    "group_music_concurrency_limit": 123,
    "character_count": 123,
    "is_scim_synced": true,
    "scim_group": {
      "scim_external_id": "value",
      "display_name": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "seat_type": "workspace_admin"
    },
    "scim_frozen": true
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceGroupsClient(gotResponse, []byte(`{
  "key": {
    "name": "value",
    "id": "value",
    "members": [
      "value"
    ],
    "permissions": [
      "text_to_speech"
    ],
    "group_usage_limit": 123,
    "group_pvc_limit": 123,
    "group_concurrency_limit": 123,
    "group_dubbing_concurrency_limit": 123,
    "group_music_concurrency_limit": 123,
    "character_count": 123,
    "is_scim_synced": true,
    "scim_group": {
      "scim_external_id": "value",
      "display_name": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "seat_type": "workspace_admin"
    },
    "scim_frozen": true
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": {
    "name": "value",
    "id": "value",
    "members": [
      "value"
    ],
    "permissions": [
      "text_to_speech"
    ],
    "group_usage_limit": 123,
    "group_pvc_limit": 123,
    "group_concurrency_limit": 123,
    "group_dubbing_concurrency_limit": 123,
    "group_music_concurrency_limit": 123,
    "character_count": 123,
    "is_scim_synced": true,
    "scim_group": {
      "scim_external_id": "value",
      "display_name": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "seat_type": "workspace_admin"
    },
    "scim_frozen": true
  }
}`)
	}

	req := receiveWireWorkspaceGroupsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/groups" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/groups")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceGroupsClient(t, 422, `{
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

	_, err := client.List(context.Background())
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

func TestSearchWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceGroupsClient(t, 200, `[
  {
    "name": "value",
    "id": "value",
    "members_emails": [
      "value"
    ]
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &workspace.SearchGroupsRequest{
		Name: "value",
	}
	result, err := client.Search(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceGroupsClient(gotResponse, []byte(`[
  {
    "name": "value",
    "id": "value",
    "members_emails": [
      "value"
    ]
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "name": "value",
    "id": "value",
    "members_emails": [
      "value"
    ]
  }
]`)
	}

	req := receiveWireWorkspaceGroupsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/groups/search" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/groups/search")
	}
	if got := req.Query.Get("name"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "name", got, "value")
	}
}

func TestSearchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceGroupsClient(t, 422, `{
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

	request := &workspace.SearchGroupsRequest{
		Name: "value",
	}
	_, err := client.Search(context.Background(), request)
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

type wireRequestWorkspaceGroupsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceGroupsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceGroupsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceGroupsClient) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceGroupsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceGroupsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceGroupsClient(t *testing.T, recorded <-chan wireRequestWorkspaceGroupsClient) wireRequestWorkspaceGroupsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceGroupsClient{}
	}
}

// wireJSONContainsWorkspaceGroupsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceGroupsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceGroupsClient(got, want)
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

// setWireValueWorkspaceGroupsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceGroupsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceGroupsClient(got, want []byte) bool {
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
