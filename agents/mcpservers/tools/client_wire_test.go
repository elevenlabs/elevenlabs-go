package tools

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

	mcpservers "github.com/elevenlabs/elevenlabs-go/agents/mcpservers"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversTools(t, 200, `{
  "success": true,
  "tools": [
    {
      "name": "value",
      "title": "value",
      "description": "value",
      "inputSchema": {
        "key": {}
      },
      "execution": {
        "taskSupport": "forbidden"
      },
      "outputSchema": {
        "key": {}
      },
      "icons": [
        {
          "src": "value",
          "mimeType": "value",
          "sizes": [
            "value"
          ],
          "theme": "light"
        }
      ],
      "annotations": {
        "title": "value",
        "readOnlyHint": true,
        "destructiveHint": true,
        "idempotentHint": true,
        "openWorldHint": true
      },
      "_meta": {
        "key": {}
      }
    }
  ],
  "tool_approval_statuses": [
    {
      "tool_id": "value",
      "state": "up_to_date",
      "approval_policy": "auto_approved",
      "approved_definition": {
        "description": "value",
        "input_schema": {
          "key": {}
        }
      }
    }
  ],
  "error_message": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	environmentValue := string("value")
	request := &mcpservers.ListToolsRequest{
		Environment: &environmentValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversTools(gotResponse, []byte(`{
  "success": true,
  "tools": [
    {
      "name": "value",
      "title": "value",
      "description": "value",
      "inputSchema": {
        "key": {}
      },
      "execution": {
        "taskSupport": "forbidden"
      },
      "outputSchema": {
        "key": {}
      },
      "icons": [
        {
          "src": "value",
          "mimeType": "value",
          "sizes": [
            "value"
          ],
          "theme": "light"
        }
      ],
      "annotations": {
        "title": "value",
        "readOnlyHint": true,
        "destructiveHint": true,
        "idempotentHint": true,
        "openWorldHint": true
      },
      "_meta": {
        "key": {}
      }
    }
  ],
  "tool_approval_statuses": [
    {
      "tool_id": "value",
      "state": "up_to_date",
      "approval_policy": "auto_approved",
      "approved_definition": {
        "description": "value",
        "input_schema": {
          "key": {}
        }
      }
    }
  ],
  "error_message": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "success": true,
  "tools": [
    {
      "name": "value",
      "title": "value",
      "description": "value",
      "inputSchema": {
        "key": {}
      },
      "execution": {
        "taskSupport": "forbidden"
      },
      "outputSchema": {
        "key": {}
      },
      "icons": [
        {
          "src": "value",
          "mimeType": "value",
          "sizes": [
            "value"
          ],
          "theme": "light"
        }
      ],
      "annotations": {
        "title": "value",
        "readOnlyHint": true,
        "destructiveHint": true,
        "idempotentHint": true,
        "openWorldHint": true
      },
      "_meta": {
        "key": {}
      }
    }
  ],
  "tool_approval_statuses": [
    {
      "tool_id": "value",
      "state": "up_to_date",
      "approval_policy": "auto_approved",
      "approved_definition": {
        "description": "value",
        "input_schema": {
          "key": {}
        }
      }
    }
  ],
  "error_message": "value"
}`)
	}

	req := receiveWireAgentsMcpserversTools(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/mcp-servers/value/tools" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value/tools")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversTools(t, 422, `{
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

	environmentValue := string("value")
	request := &mcpservers.ListToolsRequest{
		Environment: &environmentValue,
	}
	_, err := client.List(context.Background(), string("value"), request)
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

type wireRequestAgentsMcpserversTools struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsMcpserversTools records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsMcpserversTools(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsMcpserversTools) {
	t.Helper()
	recorded := make(chan wireRequestAgentsMcpserversTools, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsMcpserversTools{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsMcpserversTools(t *testing.T, recorded <-chan wireRequestAgentsMcpserversTools) wireRequestAgentsMcpserversTools {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsMcpserversTools{}
	}
}

// wireJSONContainsAgentsMcpserversTools reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsMcpserversTools(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsMcpserversTools(got, want)
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

// setWireValueAgentsMcpserversTools sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsMcpserversTools(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsMcpserversTools(got, want []byte) bool {
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
