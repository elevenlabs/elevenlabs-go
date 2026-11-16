package messages

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

func TestTextSearchWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsMessages(t, 200, `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	callDurationMinSecsValue := int(123)
	callDurationMaxSecsValue := int(123)
	ratingMaxValue := int(5)
	ratingMinValue := int(5)
	hasFeedbackCommentValue := bool(true)
	userIdValue := string("value")
	includeInvalidToolCallsValue := bool(true)
	pageSizeValue := int(50)
	textOnlyValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	cursorValue := string("value")
	request := &conversations.TextSearchMessagesRequest{
		TextQuery:               "value",
		AgentId:                 &agentIdValue,
		CallStartBeforeUnix:     &callStartBeforeUnixValue,
		CallStartAfterUnix:      &callStartAfterUnixValue,
		CallDurationMinSecs:     &callDurationMinSecsValue,
		CallDurationMaxSecs:     &callDurationMaxSecsValue,
		RatingMax:               &ratingMaxValue,
		RatingMin:               &ratingMinValue,
		HasFeedbackComment:      &hasFeedbackCommentValue,
		UserId:                  &userIdValue,
		IncludeInvalidToolCalls: &includeInvalidToolCallsValue,
		PageSize:                &pageSizeValue,
		TextOnly:                &textOnlyValue,
		BranchId:                &branchIdValue,
		VersionId:               &versionIdValue,
		Cursor:                  &cursorValue,
	}
	result, err := client.TextSearch(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsMessages(gotResponse, []byte(`{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsConversationsMessages(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/messages/text-search" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/messages/text-search")
	}
	if got := req.Query.Get("text_query"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "text_query", got, "value")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("call_start_before_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_start_before_unix", got, "123")
	}
	if got := req.Query.Get("call_start_after_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_start_after_unix", got, "123")
	}
	if got := req.Query.Get("call_duration_min_secs"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_duration_min_secs", got, "123")
	}
	if got := req.Query.Get("call_duration_max_secs"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "call_duration_max_secs", got, "123")
	}
	if got := req.Query.Get("rating_max"); got != "5" {
		t.Fatalf("query %s: got %q want %q", "rating_max", got, "5")
	}
	if got := req.Query.Get("rating_min"); got != "5" {
		t.Fatalf("query %s: got %q want %q", "rating_min", got, "5")
	}
	if got := req.Query.Get("has_feedback_comment"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "has_feedback_comment", got, "true")
	}
	if got := req.Query.Get("user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "user_id", got, "value")
	}
	if got := req.Query.Get("include_invalid_tool_calls"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_invalid_tool_calls", got, "true")
	}
	if got := req.Query.Get("page_size"); got != "50" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "50")
	}
	if got := req.Query.Get("text_only"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "text_only", got, "true")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestTextSearchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsMessages(t, 422, `{
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
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	callDurationMinSecsValue := int(123)
	callDurationMaxSecsValue := int(123)
	ratingMaxValue := int(5)
	ratingMinValue := int(5)
	hasFeedbackCommentValue := bool(true)
	userIdValue := string("value")
	includeInvalidToolCallsValue := bool(true)
	pageSizeValue := int(50)
	textOnlyValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	cursorValue := string("value")
	request := &conversations.TextSearchMessagesRequest{
		TextQuery:               "value",
		AgentId:                 &agentIdValue,
		CallStartBeforeUnix:     &callStartBeforeUnixValue,
		CallStartAfterUnix:      &callStartAfterUnixValue,
		CallDurationMinSecs:     &callDurationMinSecsValue,
		CallDurationMaxSecs:     &callDurationMaxSecsValue,
		RatingMax:               &ratingMaxValue,
		RatingMin:               &ratingMinValue,
		HasFeedbackComment:      &hasFeedbackCommentValue,
		UserId:                  &userIdValue,
		IncludeInvalidToolCalls: &includeInvalidToolCallsValue,
		PageSize:                &pageSizeValue,
		TextOnly:                &textOnlyValue,
		BranchId:                &branchIdValue,
		VersionId:               &versionIdValue,
		Cursor:                  &cursorValue,
	}
	_, err := client.TextSearch(context.Background(), request)
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

	server, recorded := serveWireAgentsConversationsMessages(t, 200, `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
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
	pageSizeValue := int(50)
	cursorValue := string("value")
	request := &conversations.SearchMessagesRequest{
		TextQuery: "value",
		AgentId:   &agentIdValue,
		BranchId:  &branchIdValue,
		PageSize:  &pageSizeValue,
		Cursor:    &cursorValue,
	}
	result, err := client.Search(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsMessages(gotResponse, []byte(`{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "conversation_id": "value",
      "agent_id": "value",
      "agent_name": "value",
      "transcript_index": 123,
      "chunk_text": "value",
      "chunk_highlights": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99,
      "conversation_start_time_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsConversationsMessages(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/messages/smart-search" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/messages/smart-search")
	}
	if got := req.Query.Get("text_query"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "text_query", got, "value")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "50" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "50")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestSearchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsMessages(t, 422, `{
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
	pageSizeValue := int(50)
	cursorValue := string("value")
	request := &conversations.SearchMessagesRequest{
		TextQuery: "value",
		AgentId:   &agentIdValue,
		BranchId:  &branchIdValue,
		PageSize:  &pageSizeValue,
		Cursor:    &cursorValue,
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

type wireRequestAgentsConversationsMessages struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsConversationsMessages records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsConversationsMessages(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsConversationsMessages) {
	t.Helper()
	recorded := make(chan wireRequestAgentsConversationsMessages, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsConversationsMessages{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsConversationsMessages(t *testing.T, recorded <-chan wireRequestAgentsConversationsMessages) wireRequestAgentsConversationsMessages {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsConversationsMessages{}
	}
}

// wireJSONContainsAgentsConversationsMessages reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsConversationsMessages(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsConversationsMessages(got, want)
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

// setWireValueAgentsConversationsMessages sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsConversationsMessages(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsConversationsMessages(got, want []byte) bool {
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
