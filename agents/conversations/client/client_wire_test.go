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

func TestGetSignedUrlWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "signed_url": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	includeConversationIdValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	environmentValue := string("value")
	debugEventsRequestValue := bool(true)
	request := &agents.GetSignedUrlConversationsRequest{
		AgentId:               "value",
		IncludeConversationId: &includeConversationIdValue,
		BranchId:              &branchIdValue,
		VersionId:             &versionIdValue,
		Environment:           &environmentValue,
		DebugEventsRequest:    &debugEventsRequestValue,
	}
	result, err := client.GetSignedUrl(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "signed_url": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "signed_url": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversation/get-signed-url" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversation/get-signed-url")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("include_conversation_id"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_conversation_id", got, "true")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
	if got := req.Query.Get("debug_events_request"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "debug_events_request", got, "true")
	}
}

func TestGetSignedUrlWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	includeConversationIdValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	environmentValue := string("value")
	debugEventsRequestValue := bool(true)
	request := &agents.GetSignedUrlConversationsRequest{
		AgentId:               "value",
		IncludeConversationId: &includeConversationIdValue,
		BranchId:              &branchIdValue,
		VersionId:             &versionIdValue,
		Environment:           &environmentValue,
		DebugEventsRequest:    &debugEventsRequestValue,
	}
	_, err := client.GetSignedUrl(context.Background(), request)
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

func TestGetWebrtcTokenWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "token": "value",
  "conversation_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	participantNameValue := string("value")
	branchIdValue := string("value")
	versionIdValue := string("value")
	environmentValue := string("value")
	debugEventsRequestValue := bool(true)
	request := &agents.GetWebrtcTokenConversationsRequest{
		AgentId:            "value",
		ParticipantName:    &participantNameValue,
		BranchId:           &branchIdValue,
		VersionId:          &versionIdValue,
		Environment:        &environmentValue,
		DebugEventsRequest: &debugEventsRequestValue,
	}
	result, err := client.GetWebrtcToken(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "token": "value",
  "conversation_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "token": "value",
  "conversation_id": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversation/token" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversation/token")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("participant_name"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "participant_name", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
	if got := req.Query.Get("debug_events_request"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "debug_events_request", got, "true")
	}
}

func TestGetWebrtcTokenWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	participantNameValue := string("value")
	branchIdValue := string("value")
	versionIdValue := string("value")
	environmentValue := string("value")
	debugEventsRequestValue := bool(true)
	request := &agents.GetWebrtcTokenConversationsRequest{
		AgentId:            "value",
		ParticipantName:    &participantNameValue,
		BranchId:           &branchIdValue,
		VersionId:          &versionIdValue,
		Environment:        &environmentValue,
		DebugEventsRequest: &debugEventsRequestValue,
	}
	_, err := client.GetWebrtcToken(context.Background(), request)
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "conversations": [
    {
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "agent_name": "value",
      "conversation_id": "value",
      "start_time_unix_secs": 123,
      "call_duration_secs": 123,
      "cost": 123,
      "cost_fiat": 9.99,
      "message_count": 123,
      "user_message_count": 123,
      "agent_message_count": 123,
      "status": "initiated",
      "termination_reason": "value",
      "call_successful": "success",
      "call_success_score": 9.99,
      "transcript_summary": "value",
      "call_summary_title": "value",
      "main_language": "value",
      "conversation_initiation_source": "unknown",
      "tool_names": [
        "value"
      ],
      "direction": "inbound",
      "rating": 9.99,
      "sentiment_analysis": {
        "overall_label": "positive",
        "overall_sentiment_score": 0.5,
        "overall_frustration_score": 0.5,
        "min_user_sentiment_score": 0.5,
        "max_user_frustration_score": 0.5,
        "num_scored_user_turns": 123
      },
      "data_collection_results": {
        "key": {}
      },
      "evaluation_criteria_results": {
        "key": {
          "result": "success",
          "score": 123,
          "max_score": 123
        }
      },
      "tag_ids": [
        "value"
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	agentIdValue := string("value")
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	callDurationMinSecsValue := int(123)
	callDurationMaxSecsValue := int(123)
	ratingMaxValue := int(5)
	ratingMinValue := int(5)
	hasFeedbackCommentValue := bool(true)
	userIdValue := string("value")
	dataCollectionFilterValue := string("value")
	dynamicVariableFilterValue := string("value")
	includeInvalidToolCallsValue := bool(true)
	pageSizeValue := int(100)
	searchValue := string("value")
	textOnlyValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	parentConversationIdValue := string("value")
	workflowNodeEnteredIdValue := string("value")
	request := &agents.ListConversationsRequest{
		Cursor:                  &cursorValue,
		AgentId:                 &agentIdValue,
		CallStartBeforeUnix:     &callStartBeforeUnixValue,
		CallStartAfterUnix:      &callStartAfterUnixValue,
		CallDurationMinSecs:     &callDurationMinSecsValue,
		CallDurationMaxSecs:     &callDurationMaxSecsValue,
		RatingMax:               &ratingMaxValue,
		RatingMin:               &ratingMinValue,
		HasFeedbackComment:      &hasFeedbackCommentValue,
		UserId:                  &userIdValue,
		DataCollectionFilter:    &dataCollectionFilterValue,
		DynamicVariableFilter:   &dynamicVariableFilterValue,
		IncludeInvalidToolCalls: &includeInvalidToolCallsValue,
		PageSize:                &pageSizeValue,
		Search:                  &searchValue,
		TextOnly:                &textOnlyValue,
		BranchId:                &branchIdValue,
		VersionId:               &versionIdValue,
		ParentConversationId:    &parentConversationIdValue,
		WorkflowNodeEnteredId:   &workflowNodeEnteredIdValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "conversations": [
    {
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "agent_name": "value",
      "conversation_id": "value",
      "start_time_unix_secs": 123,
      "call_duration_secs": 123,
      "cost": 123,
      "cost_fiat": 9.99,
      "message_count": 123,
      "user_message_count": 123,
      "agent_message_count": 123,
      "status": "initiated",
      "termination_reason": "value",
      "call_successful": "success",
      "call_success_score": 9.99,
      "transcript_summary": "value",
      "call_summary_title": "value",
      "main_language": "value",
      "conversation_initiation_source": "unknown",
      "tool_names": [
        "value"
      ],
      "direction": "inbound",
      "rating": 9.99,
      "sentiment_analysis": {
        "overall_label": "positive",
        "overall_sentiment_score": 0.5,
        "overall_frustration_score": 0.5,
        "min_user_sentiment_score": 0.5,
        "max_user_frustration_score": 0.5,
        "num_scored_user_turns": 123
      },
      "data_collection_results": {
        "key": {}
      },
      "evaluation_criteria_results": {
        "key": {
          "result": "success",
          "score": 123,
          "max_score": 123
        }
      },
      "tag_ids": [
        "value"
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversations": [
    {
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "agent_name": "value",
      "conversation_id": "value",
      "start_time_unix_secs": 123,
      "call_duration_secs": 123,
      "cost": 123,
      "cost_fiat": 9.99,
      "message_count": 123,
      "user_message_count": 123,
      "agent_message_count": 123,
      "status": "initiated",
      "termination_reason": "value",
      "call_successful": "success",
      "call_success_score": 9.99,
      "transcript_summary": "value",
      "call_summary_title": "value",
      "main_language": "value",
      "conversation_initiation_source": "unknown",
      "tool_names": [
        "value"
      ],
      "direction": "inbound",
      "rating": 9.99,
      "sentiment_analysis": {
        "overall_label": "positive",
        "overall_sentiment_score": 0.5,
        "overall_frustration_score": 0.5,
        "min_user_sentiment_score": 0.5,
        "max_user_frustration_score": 0.5,
        "num_scored_user_turns": 123
      },
      "data_collection_results": {
        "key": {}
      },
      "evaluation_criteria_results": {
        "key": {
          "result": "success",
          "score": 123,
          "max_score": 123
        }
      },
      "tag_ids": [
        "value"
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
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
	if got := req.Query.Get("data_collection_filter"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "data_collection_filter", got, "value")
	}
	if got := req.Query.Get("dynamic_variable_filter"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "dynamic_variable_filter", got, "value")
	}
	if got := req.Query.Get("include_invalid_tool_calls"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_invalid_tool_calls", got, "true")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
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
	if got := req.Query.Get("parent_conversation_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "parent_conversation_id", got, "value")
	}
	if got := req.Query.Get("workflow_node_entered_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "workflow_node_entered_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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
	agentIdValue := string("value")
	callStartBeforeUnixValue := int(123)
	callStartAfterUnixValue := int(123)
	callDurationMinSecsValue := int(123)
	callDurationMaxSecsValue := int(123)
	ratingMaxValue := int(5)
	ratingMinValue := int(5)
	hasFeedbackCommentValue := bool(true)
	userIdValue := string("value")
	dataCollectionFilterValue := string("value")
	dynamicVariableFilterValue := string("value")
	includeInvalidToolCallsValue := bool(true)
	pageSizeValue := int(100)
	searchValue := string("value")
	textOnlyValue := bool(true)
	branchIdValue := string("value")
	versionIdValue := string("value")
	parentConversationIdValue := string("value")
	workflowNodeEnteredIdValue := string("value")
	request := &agents.ListConversationsRequest{
		Cursor:                  &cursorValue,
		AgentId:                 &agentIdValue,
		CallStartBeforeUnix:     &callStartBeforeUnixValue,
		CallStartAfterUnix:      &callStartAfterUnixValue,
		CallDurationMinSecs:     &callDurationMinSecsValue,
		CallDurationMaxSecs:     &callDurationMaxSecsValue,
		RatingMax:               &ratingMaxValue,
		RatingMin:               &ratingMinValue,
		HasFeedbackComment:      &hasFeedbackCommentValue,
		UserId:                  &userIdValue,
		DataCollectionFilter:    &dataCollectionFilterValue,
		DynamicVariableFilter:   &dynamicVariableFilterValue,
		IncludeInvalidToolCalls: &includeInvalidToolCallsValue,
		PageSize:                &pageSizeValue,
		Search:                  &searchValue,
		TextOnly:                &textOnlyValue,
		BranchId:                &branchIdValue,
		VersionId:               &versionIdValue,
		ParentConversationId:    &parentConversationIdValue,
		WorkflowNodeEnteredId:   &workflowNodeEnteredIdValue,
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

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "conversation_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateAgentConversationRequest
	if err := json.Unmarshal([]byte(`{
  "agent": {
    "type": "persisted",
    "agent_id": "value",
    "branch_id": "value",
    "overrides": {
      "prompt": "value",
      "llm": "gpt-4o-mini",
      "language": "value",
      "tool_ids": [
        "value"
      ],
      "native_mcp_server_ids": [
        "value"
      ],
      "knowledge_base_ids": [
        "value"
      ],
      "procedure_ids": [
        "value"
      ],
      "custom_llm_extra_body": {
        "key": {}
      }
    }
  },
  "environment": "value",
  "user_id": "value",
  "dynamic_variables": {
    "key": {}
  }
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
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "conversation_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_id": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/conversations" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsConversationsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	var request agents.CreateAgentConversationRequest
	if err := json.Unmarshal([]byte(`{
  "agent": {
    "type": "persisted",
    "agent_id": "value",
    "branch_id": "value",
    "overrides": {
      "prompt": "value",
      "llm": "gpt-4o-mini",
      "language": "value",
      "tool_ids": [
        "value"
      ],
      "native_mcp_server_ids": [
        "value"
      ],
      "knowledge_base_ids": [
        "value"
      ],
      "procedure_ids": [
        "value"
      ],
      "custom_llm_extra_body": {
        "key": {}
      }
    }
  },
  "environment": "value",
  "user_id": "value",
  "dynamic_variables": {
    "key": {}
  }
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

func TestResolveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &agents.ResolveConversationsRequest{
		AgentId:   "value",
		Reference: "value",
	}
	result, err := client.Resolve(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/resolve" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/resolve")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("reference"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "reference", got, "value")
	}
}

func TestResolveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	request := &agents.ResolveConversationsRequest{
		AgentId:   "value",
		Reference: "value",
	}
	_, err := client.Resolve(context.Background(), request)
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

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &agents.GetConversationsRequest{}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_id": "value",
  "agent_name": "value",
  "conversation_product": "value",
  "status": "initiated",
  "user_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "metadata": {
    "start_time_unix_secs": 123,
    "accepted_time_unix_secs": 123,
    "call_duration_secs": 123,
    "queue_wait_secs": 9.99,
    "cost": 123,
    "deletion_settings": {
      "deletion_time_unix_secs": 123,
      "deleted_logs_at_time_unix_secs": 123,
      "deleted_audio_at_time_unix_secs": 123,
      "deleted_transcript_at_time_unix_secs": 123,
      "delete_transcript_and_pii": true,
      "delete_audio": true
    },
    "feedback": {
      "type": "thumbs",
      "overall_score": "like",
      "likes": 123,
      "dislikes": 123,
      "rating": 123,
      "comment": "value"
    },
    "authorization_method": "invalid",
    "charging": {
      "dev_discount": true,
      "is_burst": true,
      "tier": "value",
      "llm_usage": {
        "irreversible_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "initiated_generation": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        }
      },
      "llm_price": 9.99,
      "llm_charge": 123,
      "call_charge": 123,
      "platform_charge": 123,
      "platform_usage": {
        "category_usage": {
          "key": {
            "credits": 123,
            "price": 9.99,
            "quantity": 9.99
          }
        }
      },
      "platform_price": 9.99,
      "free_minutes_consumed": 9.99,
      "free_llm_dollars_consumed": 9.99,
      "tts_usage": {
        "primary_tts_model": "value",
        "total_audio_output_seconds": 9.99,
        "total_characters": 123,
        "per_voice_usage": [
          {
            "voice_id": "value",
            "audio_output_seconds": 9.99
          }
        ]
      },
      "asr_usage": {
        "asr_model": "value",
        "total_transcription_calls": 123,
        "total_audio_input_seconds": 9.99
      },
      "analysis": {
        "total": {
          "price": 9.99,
          "charge": 123,
          "runs": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        },
        "last_run": {
          "price": 9.99,
          "charge": 123,
          "price_per_feature": {
            "key": 9.99
          },
          "charge_per_feature": {
            "key": 123
          }
        }
      },
      "voice_surcharge_charge": 123
    },
    "phone_call": {
      "direction": "inbound",
      "phone_number_id": "value",
      "agent_number": "value",
      "external_number": "value",
      "call_result": "answered",
      "type": "twilio",
      "stream_sid": "value",
      "call_sid": "value"
    },
    "batch_call": {
      "batch_call_id": "value",
      "batch_call_recipient_id": "value",
      "campaign": {
        "campaign_id": "value",
        "campaign_lead_id": "value",
        "cold_attempt": 123,
        "conversation_type": "initial"
      }
    },
    "termination_reason": "value",
    "error": {
      "code": 123,
      "reason": "value"
    },
    "warnings": [
      "value"
    ],
    "main_language": "value",
    "rag_usage": {
      "usage_count": 123,
      "embedding_model": "value"
    },
    "text_only": true,
    "features_usage": {
      "language_detection": {
        "enabled": true,
        "used": true
      },
      "transfer_to_agent": {
        "enabled": true,
        "used": true
      },
      "transfer_to_number": {
        "enabled": true,
        "used": true
      },
      "multivoice": {
        "enabled": true,
        "used": true
      },
      "dtmf_tones": {
        "enabled": true,
        "used": true
      },
      "external_mcp_servers": {
        "enabled": true,
        "used": true
      },
      "pii_zrm_workspace": true,
      "pii_zrm_agent": true,
      "tool_dynamic_variable_updates": {
        "enabled": true,
        "used": true
      },
      "is_livekit": true,
      "voicemail_detection": {
        "enabled": true,
        "used": true
      },
      "dtmf_input": {
        "enabled": true,
        "used": true
      },
      "workflow": {
        "enabled": true,
        "tool_node": {
          "enabled": true,
          "used": true
        },
        "standalone_agent_node": {
          "enabled": true,
          "used": true
        },
        "phone_number_node": {
          "enabled": true,
          "used": true
        },
        "end_node": {
          "enabled": true,
          "used": true
        },
        "override_agent_node": {
          "enabled": true,
          "used": true
        }
      },
      "agent_testing": {
        "enabled": true,
        "tests_ran_after_last_modification": true,
        "tests_ran_in_last_7_days": true
      },
      "versioning": {
        "enabled": true,
        "used": true
      },
      "file_input": {
        "enabled": true,
        "used": true
      },
      "freeform_procedure": {
        "enabled": true,
        "used": true
      },
      "structured_procedure": {
        "enabled": true,
        "used": true
      }
    },
    "eleven_assistant": {
      "is_eleven_assistant": true
    },
    "initiator_id": "value",
    "conversation_initiation_source": "unknown",
    "conversation_initiation_source_version": "value",
    "timezone": "value",
    "async_metadata": {
      "delivery_status": "pending",
      "delivery_timestamp": 123,
      "delivery_error": "value",
      "external_system": "value",
      "external_id": "value",
      "external_link": "value",
      "retry_count": 123,
      "last_retry_timestamp": 123,
      "last_processed_external_message_id": "value",
      "integration_trigger_id": "value",
      "platform_metadata": {
        "conversation_initiation_source": "plain_integration",
        "thread_id": "value",
        "workspace_id": "value"
      }
    },
    "whatsapp": {
      "direction": "inbound",
      "whatsapp_phone_number_id": "value",
      "whatsapp_user_id": "value",
      "awaiting_first_user_message": true,
      "recipient": {
        "phone": "value",
        "bsuid": "value"
      },
      "call_permission_request": {
        "state": "blocked",
        "requested_at_unix": 123,
        "responded_at_unix": 123,
        "permanent": true,
        "granted_by_user_call": true
      },
      "call_result": "answered",
      "opening_message": {
        "template_name": "value",
        "wamid": "value",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_code": "value",
        "pricing_category": "value",
        "billable": true
      }
    },
    "sms": {
      "direction": "inbound",
      "phone_number_id": "value",
      "sms_user_phone_number": "value",
      "agent_phone_number": "value"
    },
    "agent_created_from": "cli",
    "agent_last_updated_from": "cli",
    "voice_rewards": [
      {
        "voice_id": "value",
        "reward_usd_cents": 9.99
      }
    ],
    "cost_fiat": 9.99
  },
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  },
  "user_message_count": 123,
  "agent_message_count": 123,
  "visited_agents": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "conversation_initiation_client_data": {
    "conversation_config_override": {
      "asr": {
        "keywords": [
          "value"
        ]
      },
      "turn": {
        "soft_timeout_config": {
          "message": "value",
          "additional_soft_timeout_messages": [
            "value"
          ]
        },
        "noisy_environment_message": "value"
      },
      "tts": {
        "model_id": "eleven_turbo_v2",
        "voice_id": "value",
        "supported_voices": [
          {
            "label": "value",
            "voice_id": "value",
            "description": "value",
            "language": "value",
            "model_family": "turbo",
            "optimize_streaming_latency": 123,
            "stability": 0.5,
            "speed": 1.2,
            "similarity_boost": 0.5,
            "clone_strength": 0.5
          }
        ],
        "stability": 9.99,
        "speed": 9.99,
        "similarity_boost": 9.99,
        "clone_strength": 9.99,
        "pronunciation_dictionary_locators": [
          {
            "pronunciation_dictionary_id": "value",
            "version_id": "value"
          }
        ]
      },
      "conversation": {
        "text_only": true,
        "max_duration_seconds": 123
      },
      "agent": {
        "first_message": "value",
        "language": "value",
        "max_conversation_duration_message": "value",
        "prompt": {
          "prompt": "value",
          "llm": "gpt-4o-mini",
          "reasoning_effort": "none",
          "tool_ids": [
            "value"
          ],
          "tool_ids_strict": true,
          "native_mcp_server_ids": [
            "value"
          ],
          "knowledge_base": [
            {
              "type": "file",
              "name": "value",
              "id": "value",
              "usage_mode": "prompt"
            }
          ]
        }
      }
    },
    "custom_llm_extra_body": {
      "key": {}
    },
    "user_id": "value",
    "source_info": {
      "source": "unknown",
      "version": "value"
    },
    "branch_id": "value",
    "environment": "value",
    "starting_workflow_node_id": "value",
    "procedure_ids": [
      "value"
    ],
    "start_procedure_ids": [
      "value"
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "environment": "value",
  "conversation_id": "value",
  "has_audio": true,
  "has_user_audio": true,
  "has_response_audio": true,
  "has_auxiliary_audio": true,
  "transcript": [
    {
      "role": "user",
      "agent_metadata": {
        "agent_id": "value",
        "branch_id": "value",
        "workflow_node_id": "value",
        "version_id": "value"
      },
      "message": "value",
      "multivoice_message": {
        "parts": [
          {
            "text": "value",
            "voice_label": "value",
            "time_in_call_secs": 123
          }
        ]
      },
      "tool_calls": [
        {
          "type": "system",
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "params_as_json": "value",
          "tool_has_been_called": true,
          "tool_details": {
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
      "tool_results": [
        {
          "request_id": "value",
          "tool_name": "value",
          "tool_id": "value",
          "result_value": "value",
          "is_error": true,
          "is_blocked": true,
          "tool_has_been_called": true,
          "tool_latency_secs": 9.99,
          "error_type": "value",
          "raw_error_message": "value",
          "dynamic_variable_updates": [
            {
              "variable_name": "value",
              "old_value": "value",
              "new_value": "value",
              "updated_at": 9.99,
              "tool_name": "value",
              "tool_request_id": "value"
            }
          ],
          "type": "client"
        }
      ],
      "feedback": {
        "score": "like",
        "time_in_call_secs": 123
      },
      "llm_override": "value",
      "producing_llm": "value",
      "time_in_call_secs": 123,
      "conversation_turn_metrics": {
        "metrics": {
          "key": {
            "elapsed_time": 9.99
          }
        },
        "convai_asr_provider": "value",
        "convai_tts_model": "value",
        "convai_tts_cascade": "value",
        "voice_id": "value"
      },
      "rag_retrieval_info": {
        "chunks": [
          {
            "document_id": "value",
            "chunk_id": "value",
            "vector_distance": 9.99
          }
        ],
        "embedding_model": "e5_mistral_7b_instruct",
        "retrieval_query": "value",
        "rag_latency_secs": 9.99,
        "used_chunk_ids": [
          "value"
        ]
      },
      "llm_usage": {
        "model_usage": {
          "key": {
            "input": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_read": {
              "tokens": 123,
              "price": 9.99
            },
            "input_cache_write": {
              "tokens": 123,
              "price": 9.99
            },
            "output_total": {
              "tokens": 123,
              "price": 9.99
            }
          }
        }
      },
      "interrupted": true,
      "ignored_as_backchannel": true,
      "platform_event": "user_turn_timeout",
      "original_message": "value",
      "reasoning": [
        {
          "summary": "value",
          "provider_redact": true
        }
      ],
      "source_medium": "audio",
      "source_event_id": 123,
      "used_static_kb_document_ids": [
        "value"
      ],
      "user_identifier": "value",
      "id": "value",
      "created_at_unix_ms": 123,
      "triggered_guardrails": [
        {
          "guardrail_type": "custom",
          "guardrail_name": "value",
          "execution_mode": "streaming",
          "guardrail_details": {
            "type": "llm",
            "rationale": "value"
          }
        }
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "tag_ids": [
    "value"
  ],
  "otlp_traces": {
    "key": {}
  }
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	request := &agents.GetConversationsRequest{}
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{}`)

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
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/conversations/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

func TestGetSummaryWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "conversation_id": "value",
  "agent_id": "value",
  "status": "initiated",
  "call_summary_title": "value",
  "transcript_summary": "value",
  "call_successful": "success",
  "message_count": 123,
  "messages": [
    {
      "role": "user",
      "message": "value"
    }
  ],
  "messages_omitted": true,
  "note": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	maxMessagesValue := int(123)
	request := &agents.GetSummaryConversationsRequest{
		MaxMessages: &maxMessagesValue,
	}
	result, err := client.GetSummary(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "conversation_id": "value",
  "agent_id": "value",
  "status": "initiated",
  "call_summary_title": "value",
  "transcript_summary": "value",
  "call_successful": "success",
  "message_count": 123,
  "messages": [
    {
      "role": "user",
      "message": "value"
    }
  ],
  "messages_omitted": true,
  "note": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_id": "value",
  "agent_id": "value",
  "status": "initiated",
  "call_summary_title": "value",
  "transcript_summary": "value",
  "call_successful": "success",
  "message_count": 123,
  "messages": [
    {
      "role": "user",
      "message": "value"
    }
  ],
  "messages_omitted": true,
  "note": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/value/summary" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/summary")
	}
	if got := req.Query.Get("max_messages"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "max_messages", got, "123")
	}
}

func TestGetSummaryWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	maxMessagesValue := int(123)
	request := &agents.GetSummaryConversationsRequest{
		MaxMessages: &maxMessagesValue,
	}
	_, err := client.GetSummary(context.Background(), string("value"), request)
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

func TestGetSipMessagesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(20)
	cursorValue := string("value")
	request := &agents.GetSipMessagesConversationsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.GetSipMessages(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/value/sip-messages" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/sip-messages")
	}
	if got := req.Query.Get("page_size"); got != "20" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "20")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestGetSipMessagesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	pageSizeValue := int(20)
	cursorValue := string("value")
	request := &agents.GetSipMessagesConversationsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	_, err := client.GetSipMessages(context.Background(), string("value"), request)
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

func TestSendMessageWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.SendConversationMessageRequest
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "input": [
    {
      "type": "user_message",
      "content": [
        {
          "type": "text",
          "text": "value"
        }
      ]
    }
  ],
  "background": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.SendMessage(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/conversations/value/messages" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/messages")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsConversationsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSendMessageWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	var request agents.SendConversationMessageRequest
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "input": [
    {
      "type": "user_message",
      "content": [
        {
          "type": "text",
          "text": "value"
        }
      ]
    }
  ],
  "background": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.SendMessage(context.Background(), string("value"), &request)
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

func TestGetTurnWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsClient(t, 200, `{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetTurn(context.Background(), string("value"), string("10"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsClient(gotResponse, []byte(`{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`)
	}

	req := receiveWireAgentsConversationsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/conversations/value/turns/10" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/conversations/value/turns/10")
	}
}

func TestGetTurnWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsClient(t, 422, `{
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

	_, err := client.GetTurn(context.Background(), string("value"), string("10"))
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

type wireRequestAgentsConversationsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsConversationsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsConversationsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsConversationsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsConversationsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsConversationsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsConversationsClient(t *testing.T, recorded <-chan wireRequestAgentsConversationsClient) wireRequestAgentsConversationsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsConversationsClient{}
	}
}

// wireJSONContainsAgentsConversationsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsConversationsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsConversationsClient(got, want)
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

// setWireValueAgentsConversationsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsConversationsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsConversationsClient(got, want []byte) bool {
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
