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

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsProceduresClient(t, 200, `{
  "procedures": [
    {
      "procedure_id": "value",
      "version_id": "value",
      "name": "value",
      "type": "free_form",
      "trigger": "value",
      "has_draft": true,
      "referenced_tool_ids": [
        "value"
      ],
      "referenced_kb_ids": [
        "value"
      ],
      "referenced_procedure_ids": [
        "value"
      ],
      "referenced_dynamic_variables": [
        "value"
      ],
      "written_dynamic_variables": [
        "value"
      ],
      "folder_parent_id": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentVersionIdValue := string("value")
	request := &agents.ListProceduresRequest{
		AgentVersionId: &agentVersionIdValue,
	}
	result, err := client.List(context.Background(), string("value"), string("value2"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsProceduresClient(gotResponse, []byte(`{
  "procedures": [
    {
      "procedure_id": "value",
      "version_id": "value",
      "name": "value",
      "type": "free_form",
      "trigger": "value",
      "has_draft": true,
      "referenced_tool_ids": [
        "value"
      ],
      "referenced_kb_ids": [
        "value"
      ],
      "referenced_procedure_ids": [
        "value"
      ],
      "referenced_dynamic_variables": [
        "value"
      ],
      "written_dynamic_variables": [
        "value"
      ],
      "folder_parent_id": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "procedures": [
    {
      "procedure_id": "value",
      "version_id": "value",
      "name": "value",
      "type": "free_form",
      "trigger": "value",
      "has_draft": true,
      "referenced_tool_ids": [
        "value"
      ],
      "referenced_kb_ids": [
        "value"
      ],
      "referenced_procedure_ids": [
        "value"
      ],
      "referenced_dynamic_variables": [
        "value"
      ],
      "written_dynamic_variables": [
        "value"
      ],
      "folder_parent_id": "value"
    }
  ]
}`)
	}

	req := receiveWireAgentsProceduresClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/branches/value2/procedures" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/branches/value2/procedures")
	}
	if got := req.Query.Get("agent_version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_version_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsProceduresClient(t, 422, `{
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

	agentVersionIdValue := string("value")
	request := &agents.ListProceduresRequest{
		AgentVersionId: &agentVersionIdValue,
	}
	_, err := client.List(context.Background(), string("value"), string("value2"), request)
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

func TestCompileWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsProceduresClient(t, 200, `{
  "workflow": {
    "edges": {
      "key": {
        "source": "value",
        "target": "value",
        "forward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "backward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "reroutes": [
          {
            "x": 9.99,
            "y": 9.99
          }
        ]
      }
    },
    "nodes": {
      "key": {
        "type": "start",
        "position": {
          "x": 9.99,
          "y": 9.99
        },
        "edge_order": [
          "value"
        ]
      }
    },
    "prevent_subagent_loops": true
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Compile(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsProceduresClient(gotResponse, []byte(`{
  "workflow": {
    "edges": {
      "key": {
        "source": "value",
        "target": "value",
        "forward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "backward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "reroutes": [
          {
            "x": 9.99,
            "y": 9.99
          }
        ]
      }
    },
    "nodes": {
      "key": {
        "type": "start",
        "position": {
          "x": 9.99,
          "y": 9.99
        },
        "edge_order": [
          "value"
        ]
      }
    },
    "prevent_subagent_loops": true
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "workflow": {
    "edges": {
      "key": {
        "source": "value",
        "target": "value",
        "forward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "backward_condition": {
          "label": "value",
          "type": "unconditional"
        },
        "reroutes": [
          {
            "x": 9.99,
            "y": 9.99
          }
        ]
      }
    },
    "nodes": {
      "key": {
        "type": "start",
        "position": {
          "x": 9.99,
          "y": 9.99
        },
        "edge_order": [
          "value"
        ]
      }
    },
    "prevent_subagent_loops": true
  }
}`)
	}

	req := receiveWireAgentsProceduresClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/branches/value2/procedures/compile" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/branches/value2/procedures/compile")
	}
}

func TestCompileWireError400(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsProceduresClient(t, 400, `{
  "errors": {
    "key": [
      {
        "path": "value",
        "message": "value"
      }
    ]
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Compile(context.Background(), string("value"), string("value2"))
	if err == nil {
		t.Fatal("expected an error for a 400 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 400 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 400)
	}
}

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsProceduresClient(t, 200, `{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	versionIdValue := string("value")
	agentVersionIdValue := string("value")
	request := &agents.GetProceduresRequest{
		VersionId:      &versionIdValue,
		AgentVersionId: &agentVersionIdValue,
	}
	result, err := client.Get(context.Background(), string("value"), string("value2"), string("value3"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsProceduresClient(gotResponse, []byte(`{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "procedure_id": "value",
  "version_id": "value",
  "name": "value",
  "content": "value",
  "type": "free_form",
  "trigger": "value",
  "folder_parent_id": "value"
}`)
	}

	req := receiveWireAgentsProceduresClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/branches/value2/procedures/value3" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/branches/value2/procedures/value3")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
	if got := req.Query.Get("agent_version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_version_id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsProceduresClient(t, 422, `{
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

	versionIdValue := string("value")
	agentVersionIdValue := string("value")
	request := &agents.GetProceduresRequest{
		VersionId:      &versionIdValue,
		AgentVersionId: &agentVersionIdValue,
	}
	_, err := client.Get(context.Background(), string("value"), string("value2"), string("value3"), request)
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

func TestRemoveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsProceduresClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Remove(context.Background(), string("value"), string("value2"), string("value3"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsProceduresClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsProceduresClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/agents/value/branches/value2/procedures/value3" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/branches/value2/procedures/value3")
	}
}

func TestRemoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsProceduresClient(t, 422, `{
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

	_, err := client.Remove(context.Background(), string("value"), string("value2"), string("value3"))
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

type wireRequestAgentsProceduresClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsProceduresClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsProceduresClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsProceduresClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsProceduresClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsProceduresClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsProceduresClient(t *testing.T, recorded <-chan wireRequestAgentsProceduresClient) wireRequestAgentsProceduresClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsProceduresClient{}
	}
}

// wireJSONContainsAgentsProceduresClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsProceduresClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsProceduresClient(got, want)
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

// setWireValueAgentsProceduresClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsProceduresClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsProceduresClient(got, want []byte) bool {
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
