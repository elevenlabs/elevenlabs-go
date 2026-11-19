package users

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

	server, recorded := serveWireAgentsUsers(t, 200, `{
  "users": [
    {
      "user_id": "value",
      "last_contact_unix_secs": 123,
      "first_contact_unix_secs": 123,
      "conversation_count": 123,
      "last_contact_agent_id": "value",
      "last_contact_conversation_id": "value",
      "last_contact_agent_name": "value",
      "sentiment": {
        "scored_conversation_count": 123,
        "positive_count": 123,
        "neutral_count": 123,
        "negative_count": 123,
        "average_sentiment_score": 9.99,
        "average_frustration_score": 9.99,
        "recent_scored_conversation_count": 123,
        "recent_positive_count": 123,
        "recent_neutral_count": 123,
        "recent_negative_count": 123,
        "recent_average_sentiment_score": 9.99,
        "recent_average_frustration_score": 9.99
      },
      "most_frustrated_conversations": [
        {
          "conversation_id": "value",
          "agent_id": "value",
          "start_time_unix_secs": 123,
          "overall_label": "positive",
          "overall_sentiment_score": 9.99,
          "overall_frustration_score": 9.99
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	branchIdValue := string("value")
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	searchValue := string("value")
	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.ListUsersRequest{
		AgentId:             &agentIdValue,
		BranchId:            &branchIdValue,
		CallStartBeforeUnix: &callStartBeforeUnixValue,
		CallStartAfterUnix:  &callStartAfterUnixValue,
		Search:              &searchValue,
		PageSize:            &pageSizeValue,
		Cursor:              &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsUsers(gotResponse, []byte(`{
  "users": [
    {
      "user_id": "value",
      "last_contact_unix_secs": 123,
      "first_contact_unix_secs": 123,
      "conversation_count": 123,
      "last_contact_agent_id": "value",
      "last_contact_conversation_id": "value",
      "last_contact_agent_name": "value",
      "sentiment": {
        "scored_conversation_count": 123,
        "positive_count": 123,
        "neutral_count": 123,
        "negative_count": 123,
        "average_sentiment_score": 9.99,
        "average_frustration_score": 9.99,
        "recent_scored_conversation_count": 123,
        "recent_positive_count": 123,
        "recent_neutral_count": 123,
        "recent_negative_count": 123,
        "recent_average_sentiment_score": 9.99,
        "recent_average_frustration_score": 9.99
      },
      "most_frustrated_conversations": [
        {
          "conversation_id": "value",
          "agent_id": "value",
          "start_time_unix_secs": 123,
          "overall_label": "positive",
          "overall_sentiment_score": 9.99,
          "overall_frustration_score": 9.99
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "users": [
    {
      "user_id": "value",
      "last_contact_unix_secs": 123,
      "first_contact_unix_secs": 123,
      "conversation_count": 123,
      "last_contact_agent_id": "value",
      "last_contact_conversation_id": "value",
      "last_contact_agent_name": "value",
      "sentiment": {
        "scored_conversation_count": 123,
        "positive_count": 123,
        "neutral_count": 123,
        "negative_count": 123,
        "average_sentiment_score": 9.99,
        "average_frustration_score": 9.99,
        "recent_scored_conversation_count": 123,
        "recent_positive_count": 123,
        "recent_neutral_count": 123,
        "recent_negative_count": 123,
        "recent_average_sentiment_score": 9.99,
        "recent_average_frustration_score": 9.99
      },
      "most_frustrated_conversations": [
        {
          "conversation_id": "value",
          "agent_id": "value",
          "start_time_unix_secs": 123,
          "overall_label": "positive",
          "overall_sentiment_score": 9.99,
          "overall_frustration_score": 9.99
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsUsers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/users" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/users")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("call_start_before_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_start_before_unix", got, "123")
	}
	if got := req.Query.Get("call_start_after_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_start_after_unix", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
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

	server, _ := serveWireAgentsUsers(t, 422, `{
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

	agentIdValue := string("value")
	branchIdValue := string("value")
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	searchValue := string("value")
	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.ListUsersRequest{
		AgentId:             &agentIdValue,
		BranchId:            &branchIdValue,
		CallStartBeforeUnix: &callStartBeforeUnixValue,
		CallStartAfterUnix:  &callStartAfterUnixValue,
		Search:              &searchValue,
		PageSize:            &pageSizeValue,
		Cursor:              &cursorValue,
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

type wireRequestAgentsUsers struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsUsers records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsUsers(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsUsers) {
	t.Helper()
	recorded := make(chan wireRequestAgentsUsers, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsUsers{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsUsers(t *testing.T, recorded <-chan wireRequestAgentsUsers) wireRequestAgentsUsers {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsUsers{}
	}
}

// wireJSONContainsAgentsUsers reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsUsers(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsUsers(got, want)
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

// setWireValueAgentsUsers sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsUsers(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsUsers(got, want []byte) bool {
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
