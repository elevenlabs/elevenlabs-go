package executions

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

	tools "github.com/elevenlabs/elevenlabs-go/agents/tools"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsExecutions(t, 200, `{
  "executions": [
    {
      "tool_id": "value",
      "tool_request_id": "value",
      "conversation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "timestamp": 9.99,
      "latency_secs": 9.99,
      "is_error": true,
      "request_payload": "value",
      "response_payload": "value",
      "error_message": "value",
      "error_type": "value",
      "id": "value",
      "tool_call_details": {
        "type": "webhook",
        "method": "value",
        "url": "value",
        "headers": {
          "key": "value"
        },
        "path_params": {
          "key": "value"
        },
        "query_params": {
          "key": "value"
        },
        "body": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	isErrorValue := bool(true)
	agentIdValue := string("value")
	branchIdValue := string("value")
	startTimeValue := float64(9.99)
	endTimeValue := float64(9.99)
	request := &tools.GetExecutionsRequest{
		Cursor:    &cursorValue,
		PageSize:  &pageSizeValue,
		IsError:   &isErrorValue,
		AgentId:   &agentIdValue,
		BranchId:  &branchIdValue,
		StartTime: &startTimeValue,
		EndTime:   &endTimeValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsExecutions(gotResponse, []byte(`{
  "executions": [
    {
      "tool_id": "value",
      "tool_request_id": "value",
      "conversation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "timestamp": 9.99,
      "latency_secs": 9.99,
      "is_error": true,
      "request_payload": "value",
      "response_payload": "value",
      "error_message": "value",
      "error_type": "value",
      "id": "value",
      "tool_call_details": {
        "type": "webhook",
        "method": "value",
        "url": "value",
        "headers": {
          "key": "value"
        },
        "path_params": {
          "key": "value"
        },
        "query_params": {
          "key": "value"
        },
        "body": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "executions": [
    {
      "tool_id": "value",
      "tool_request_id": "value",
      "conversation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "timestamp": 9.99,
      "latency_secs": 9.99,
      "is_error": true,
      "request_payload": "value",
      "response_payload": "value",
      "error_message": "value",
      "error_type": "value",
      "id": "value",
      "tool_call_details": {
        "type": "webhook",
        "method": "value",
        "url": "value",
        "headers": {
          "key": "value"
        },
        "path_params": {
          "key": "value"
        },
        "query_params": {
          "key": "value"
        },
        "body": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsToolsExecutions(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tools/value/executions" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools/value/executions")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("is_error"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "is_error", got, "true")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("start_time"); got != "9.99" {
		t.Fatalf("query %s: got %q want %q", "start_time", got, "9.99")
	}
	if got := req.Query.Get("end_time"); got != "9.99" {
		t.Fatalf("query %s: got %q want %q", "end_time", got, "9.99")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsExecutions(t, 422, `{
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

	cursorValue := string("value")
	pageSizeValue := int(100)
	isErrorValue := bool(true)
	agentIdValue := string("value")
	branchIdValue := string("value")
	startTimeValue := float64(9.99)
	endTimeValue := float64(9.99)
	request := &tools.GetExecutionsRequest{
		Cursor:    &cursorValue,
		PageSize:  &pageSizeValue,
		IsError:   &isErrorValue,
		AgentId:   &agentIdValue,
		BranchId:  &branchIdValue,
		StartTime: &startTimeValue,
		EndTime:   &endTimeValue,
	}
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

type wireRequestAgentsToolsExecutions struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsToolsExecutions records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsToolsExecutions(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsToolsExecutions) {
	t.Helper()
	recorded := make(chan wireRequestAgentsToolsExecutions, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsToolsExecutions{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsToolsExecutions(t *testing.T, recorded <-chan wireRequestAgentsToolsExecutions) wireRequestAgentsToolsExecutions {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsToolsExecutions{}
	}
}

// wireJSONContainsAgentsToolsExecutions reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsToolsExecutions(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsToolsExecutions(got, want)
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

// setWireValueAgentsToolsExecutions sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsToolsExecutions(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsToolsExecutions(got, want []byte) bool {
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
