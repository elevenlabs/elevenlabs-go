package secrets

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

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsSecrets(t, 200, `{
  "secrets": [
    {
      "type": "stored",
      "secret_id": "value",
      "name": "value",
      "used_by": {
        "tools": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "tools_has_more": true,
        "agents": [
          {
            "referenced_resource_ids": [
              "value"
            ],
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "agents_has_more": true,
        "phone_numbers": [
          {
            "phone_number_id": "value",
            "phone_number": "value",
            "label": "value",
            "provider": "twilio"
          }
        ],
        "phone_numbers_has_more": true,
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "others": [
          "conversation_initiation_webhook"
        ]
      }
    }
  ],
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	dependencyLimitValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListSecretsRequest{
		PageSize:        &pageSizeValue,
		DependencyLimit: &dependencyLimitValue,
		Search:          &searchValue,
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
	if !wireJSONBytesEqualAgentsSecrets(gotResponse, []byte(`{
  "secrets": [
    {
      "type": "stored",
      "secret_id": "value",
      "name": "value",
      "used_by": {
        "tools": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "tools_has_more": true,
        "agents": [
          {
            "referenced_resource_ids": [
              "value"
            ],
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "agents_has_more": true,
        "phone_numbers": [
          {
            "phone_number_id": "value",
            "phone_number": "value",
            "label": "value",
            "provider": "twilio"
          }
        ],
        "phone_numbers_has_more": true,
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "others": [
          "conversation_initiation_webhook"
        ]
      }
    }
  ],
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "secrets": [
    {
      "type": "stored",
      "secret_id": "value",
      "name": "value",
      "used_by": {
        "tools": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "tools_has_more": true,
        "agents": [
          {
            "referenced_resource_ids": [
              "value"
            ],
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "agents_has_more": true,
        "phone_numbers": [
          {
            "phone_number_id": "value",
            "phone_number": "value",
            "label": "value",
            "provider": "twilio"
          }
        ],
        "phone_numbers_has_more": true,
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "others": [
          "conversation_initiation_webhook"
        ]
      }
    }
  ],
  "next_cursor": "value"
}`)
	}

	req := receiveWireAgentsSecrets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/secrets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/secrets")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("dependency_limit"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "dependency_limit", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSecrets(t, 422, `{
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
	dependencyLimitValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListSecretsRequest{
		PageSize:        &pageSizeValue,
		DependencyLimit: &dependencyLimitValue,
		Search:          &searchValue,
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsSecrets(t, 200, `{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.PostWorkspaceSecretRequest
	if err := json.Unmarshal([]byte(`{
  "type": "new",
  "name": "value",
  "value": "value"
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
	if !wireJSONBytesEqualAgentsSecrets(gotResponse, []byte(`{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)
	}

	req := receiveWireAgentsSecrets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/secrets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/secrets")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsSecrets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSecrets(t, 422, `{
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

	var request agents.PostWorkspaceSecretRequest
	if err := json.Unmarshal([]byte(`{
  "type": "new",
  "name": "value",
  "value": "value"
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

	server, recorded := serveWireAgentsSecrets(t, 200, `{
  "type": "stored",
  "secret_id": "value",
  "name": "value",
  "used_by": {
    "tools": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "tools_has_more": true,
    "agents": [
      {
        "referenced_resource_ids": [
          "value"
        ],
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "agents_has_more": true,
    "phone_numbers": [
      {
        "phone_number_id": "value",
        "phone_number": "value",
        "label": "value",
        "provider": "twilio"
      }
    ],
    "phone_numbers_has_more": true,
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "others": [
      "conversation_initiation_webhook"
    ]
  }
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
	if !wireJSONBytesEqualAgentsSecrets(gotResponse, []byte(`{
  "type": "stored",
  "secret_id": "value",
  "name": "value",
  "used_by": {
    "tools": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "tools_has_more": true,
    "agents": [
      {
        "referenced_resource_ids": [
          "value"
        ],
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "agents_has_more": true,
    "phone_numbers": [
      {
        "phone_number_id": "value",
        "phone_number": "value",
        "label": "value",
        "provider": "twilio"
      }
    ],
    "phone_numbers_has_more": true,
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "others": [
      "conversation_initiation_webhook"
    ]
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "type": "stored",
  "secret_id": "value",
  "name": "value",
  "used_by": {
    "tools": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "tools_has_more": true,
    "agents": [
      {
        "referenced_resource_ids": [
          "value"
        ],
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "agents_has_more": true,
    "phone_numbers": [
      {
        "phone_number_id": "value",
        "phone_number": "value",
        "label": "value",
        "provider": "twilio"
      }
    ],
    "phone_numbers_has_more": true,
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "others": [
      "conversation_initiation_webhook"
    ]
  }
}`)
	}

	req := receiveWireAgentsSecrets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/secrets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/secrets/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSecrets(t, 422, `{
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

	server, recorded := serveWireAgentsSecrets(t, 200, `{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.PatchWorkspaceSecretRequest
	if err := json.Unmarshal([]byte(`{
  "secret_id": "value",
  "type": "update",
  "name": "value",
  "value": "value"
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
	if !wireJSONBytesEqualAgentsSecrets(gotResponse, []byte(`{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`)
	}

	req := receiveWireAgentsSecrets(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/secrets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/secrets/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsSecrets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSecrets(t, 422, `{
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

	var request agents.PatchWorkspaceSecretRequest
	if err := json.Unmarshal([]byte(`{
  "secret_id": "value",
  "type": "update",
  "name": "value",
  "value": "value"
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

	server, recorded := serveWireAgentsSecrets(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsSecrets(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/secrets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/secrets/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSecrets(t, 422, `{
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

type wireRequestAgentsSecrets struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsSecrets records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsSecrets(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsSecrets) {
	t.Helper()
	recorded := make(chan wireRequestAgentsSecrets, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsSecrets{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsSecrets(t *testing.T, recorded <-chan wireRequestAgentsSecrets) wireRequestAgentsSecrets {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsSecrets{}
	}
}

// wireJSONContainsAgentsSecrets reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsSecrets(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsSecrets(got, want)
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

// setWireValueAgentsSecrets sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsSecrets(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsSecrets(got, want []byte) bool {
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
