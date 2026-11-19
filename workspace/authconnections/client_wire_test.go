package authconnections

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

	server, recorded := serveWireWorkspaceAuthconnections(t, 200, `{
  "auth_connections": [
    {
      "name": "value",
      "auth_type": "oauth2_client_credentials",
      "provider": "value",
      "client_id": "value",
      "token_url": "value",
      "scopes": [
        "value"
      ],
      "extra_params": {
        "key": "value"
      },
      "basic_auth_in_header": true,
      "id": "value",
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
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "integration_connections": [
          {
            "id": "value",
            "name": "value"
          }
        ]
      },
      "status": "active",
      "status_detail": "value",
      "status_updated_at": "value",
      "custom_headers": {
        "key": "value"
      }
    }
  ]
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
	if !wireJSONBytesEqualWorkspaceAuthconnections(gotResponse, []byte(`{
  "auth_connections": [
    {
      "name": "value",
      "auth_type": "oauth2_client_credentials",
      "provider": "value",
      "client_id": "value",
      "token_url": "value",
      "scopes": [
        "value"
      ],
      "extra_params": {
        "key": "value"
      },
      "basic_auth_in_header": true,
      "id": "value",
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
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "integration_connections": [
          {
            "id": "value",
            "name": "value"
          }
        ]
      },
      "status": "active",
      "status_detail": "value",
      "status_updated_at": "value",
      "custom_headers": {
        "key": "value"
      }
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "auth_connections": [
    {
      "name": "value",
      "auth_type": "oauth2_client_credentials",
      "provider": "value",
      "client_id": "value",
      "token_url": "value",
      "scopes": [
        "value"
      ],
      "extra_params": {
        "key": "value"
      },
      "basic_auth_in_header": true,
      "id": "value",
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
        "mcp_servers": [
          {
            "id": "value",
            "name": "value",
            "type": "available",
            "created_at_unix_secs": 123,
            "access_level": "admin"
          }
        ],
        "integration_connections": [
          {
            "id": "value",
            "name": "value"
          }
        ]
      },
      "status": "active",
      "status_detail": "value",
      "status_updated_at": "value",
      "custom_headers": {
        "key": "value"
      }
    }
  ]
}`)
	}

	req := receiveWireWorkspaceAuthconnections(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/auth-connections" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/auth-connections")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceAuthconnections(t, 422, `{
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceAuthconnections(t, 200, `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.CreateAuthConnectionsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "client_secret": "value",
  "custom_headers": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceAuthconnections(gotResponse, []byte(`{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)
	}

	req := receiveWireWorkspaceAuthconnections(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/auth-connections" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/auth-connections")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceAuthconnections(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceAuthconnections(t, 422, `{
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

	var request workspace.CreateAuthConnectionsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "client_secret": "value",
  "custom_headers": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceAuthconnections(t, 200, `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.UpdateAuthConnectionsRequestBody
	if err := json.Unmarshal([]byte(`{
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "client_secret": "value",
  "custom_headers": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceAuthconnections(gotResponse, []byte(`{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
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
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "integration_connections": [
      {
        "id": "value",
        "name": "value"
      }
    ]
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`)
	}

	req := receiveWireWorkspaceAuthconnections(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/workspace/auth-connections/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/auth-connections/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceAuthconnections(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceAuthconnections(t, 422, `{
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

	var request workspace.UpdateAuthConnectionsRequestBody
	if err := json.Unmarshal([]byte(`{
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "client_secret": "value",
  "custom_headers": {
    "key": "value"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
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

	server, recorded := serveWireWorkspaceAuthconnections(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceAuthconnections(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireWorkspaceAuthconnections(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/workspace/auth-connections/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/auth-connections/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceAuthconnections(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"))
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

type wireRequestWorkspaceAuthconnections struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceAuthconnections records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceAuthconnections(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceAuthconnections) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceAuthconnections, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceAuthconnections{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceAuthconnections(t *testing.T, recorded <-chan wireRequestWorkspaceAuthconnections) wireRequestWorkspaceAuthconnections {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceAuthconnections{}
	}
}

// wireJSONContainsWorkspaceAuthconnections reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceAuthconnections(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceAuthconnections(got, want)
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

// setWireValueWorkspaceAuthconnections sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceAuthconnections(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceAuthconnections(got, want []byte) bool {
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
