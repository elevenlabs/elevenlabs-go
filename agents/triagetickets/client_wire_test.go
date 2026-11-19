package triagetickets

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

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	conversationIdValue := string("value")
	ownerUserIdValue := string("value")
	assigneeUserIdValue := string("value")
	labelValue := string("value")
	mergedIntoTicketIdValue := string("value")
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListTriageTicketsRequest{
		PageSize:           &pageSizeValue,
		ConversationId:     &conversationIdValue,
		OwnerUserId:        &ownerUserIdValue,
		AssigneeUserId:     &assigneeUserIdValue,
		Label:              &labelValue,
		MergedIntoTicketId: &mergedIntoTicketIdValue,
		Search:             &searchValue,
		Cursor:             &cursorValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/triage-tickets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/triage-tickets")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("conversation_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "conversation_id", got, "value")
	}
	if got := req.Query.Get("owner_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "owner_user_id", got, "value")
	}
	if got := req.Query.Get("assignee_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "assignee_user_id", got, "value")
	}
	if got := req.Query.Get("label"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "label", got, "value")
	}
	if got := req.Query.Get("merged_into_ticket_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "merged_into_ticket_id", got, "value")
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

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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
	conversationIdValue := string("value")
	ownerUserIdValue := string("value")
	assigneeUserIdValue := string("value")
	labelValue := string("value")
	mergedIntoTicketIdValue := string("value")
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListTriageTicketsRequest{
		PageSize:           &pageSizeValue,
		ConversationId:     &conversationIdValue,
		OwnerUserId:        &ownerUserIdValue,
		AssigneeUserId:     &assigneeUserIdValue,
		Label:              &labelValue,
		MergedIntoTicketId: &mergedIntoTicketIdValue,
		Search:             &searchValue,
		Cursor:             &cursorValue,
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

func TestCreateManualWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateManualTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "title": "value",
  "qa_comment": "value",
  "priority": "low"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateManual(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/triage-tickets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/triage-tickets")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateManualWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.CreateManualTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "title": "value",
  "qa_comment": "value",
  "priority": "low"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateManual(context.Background(), string("value"), &request)
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

func TestListForWorkspaceWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	agentIdValue := string("value")
	ownerUserIdValue := string("value")
	assigneeUserIdValue := string("value")
	labelValue := string("value")
	mergedIntoTicketIdValue := string("value")
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListForWorkspaceTriageTicketsRequest{
		PageSize:           &pageSizeValue,
		AgentId:            &agentIdValue,
		OwnerUserId:        &ownerUserIdValue,
		AssigneeUserId:     &assigneeUserIdValue,
		Label:              &labelValue,
		MergedIntoTicketId: &mergedIntoTicketIdValue,
		Search:             &searchValue,
		Cursor:             &cursorValue,
	}
	result, err := client.ListForWorkspace(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/triage-tickets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("owner_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "owner_user_id", got, "value")
	}
	if got := req.Query.Get("assignee_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "assignee_user_id", got, "value")
	}
	if got := req.Query.Get("label"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "label", got, "value")
	}
	if got := req.Query.Get("merged_into_ticket_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "merged_into_ticket_id", got, "value")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListForWorkspaceWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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
	agentIdValue := string("value")
	ownerUserIdValue := string("value")
	assigneeUserIdValue := string("value")
	labelValue := string("value")
	mergedIntoTicketIdValue := string("value")
	searchValue := string("value")
	cursorValue := string("value")
	request := &agents.ListForWorkspaceTriageTicketsRequest{
		PageSize:           &pageSizeValue,
		AgentId:            &agentIdValue,
		OwnerUserId:        &ownerUserIdValue,
		AssigneeUserId:     &assigneeUserIdValue,
		Label:              &labelValue,
		MergedIntoTicketId: &mergedIntoTicketIdValue,
		Search:             &searchValue,
		Cursor:             &cursorValue,
	}
	_, err := client.ListForWorkspace(context.Background(), request)
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

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateAgentConversationTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "title": "value",
  "qa_comment": "value",
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value"
    }
  ],
  "priority": "low"
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
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/triage-tickets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.CreateAgentConversationTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "title": "value",
  "qa_comment": "value",
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value"
    }
  ],
  "priority": "low"
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

func TestListAssignableUsersWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `[
  {
    "user_id": "value",
    "email": "value",
    "first_name": "value",
    "is_service_account": true,
    "has_access": true
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.ListAssignableUsers(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`[
  {
    "user_id": "value",
    "email": "value",
    "first_name": "value",
    "is_service_account": true,
    "has_access": true
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "user_id": "value",
    "email": "value",
    "first_name": "value",
    "is_service_account": true,
    "has_access": true
  }
]`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/triage-tickets/assignable-users" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/triage-tickets/assignable-users")
	}
}

func TestListAssignableUsersWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	_, err := client.ListAssignableUsers(context.Background(), string("value"))
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

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
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
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/triage-tickets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.PatchAgentConversationTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "title": "value",
  "qa_comment": "value",
  "status": "open",
  "assignee_user_id": "value",
  "priority": "low"
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
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/triage-tickets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.PatchAgentConversationTicketRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "title": "value",
  "qa_comment": "value",
  "status": "open",
  "assignee_user_id": "value",
  "priority": "low"
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

	server, recorded := serveWireAgentsTriagetickets(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/triage-tickets/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

func TestAddCommentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.AddTicketCommentRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.AddComment(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/triage-tickets/value/comments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value/comments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAddCommentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.AddTicketCommentRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.AddComment(context.Background(), string("value"), &request)
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

func TestMergeWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.MergeAgentConversationTicketsRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "source_ticket_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Merge(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/triage-tickets/value/merge" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value/merge")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestMergeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.MergeAgentConversationTicketsRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "source_ticket_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Merge(context.Background(), string("value"), &request)
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

func TestAddTurnCommentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTriagetickets(t, 200, `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.AddTurnCommentRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "turn_index": 123,
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.AddTurnComment(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTriagetickets(gotResponse, []byte(`{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agentqa_ticket_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "agent_id": "value",
  "needs_clustering": true,
  "title": "value",
  "issue_type": "knowledge_gap",
  "labels": [
    "value"
  ],
  "conversation_ids": [
    "value"
  ],
  "first_seen_unix_secs": 123,
  "last_seen_unix_secs": 123,
  "qa_comment": "value",
  "ticket_comments": [
    {
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value",
      "created_at_unix_secs": 123,
      "owner_user_id": "value"
    }
  ],
  "status": "open",
  "priority": "low",
  "priority_changes": [
    {
      "priority": "low",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "status_changes": [
    {
      "status": "open",
      "changed_by_user_id": "value",
      "changed_at_unix_secs": 123
    }
  ],
  "merged_into_ticket_id": "value",
  "source": "qa",
  "assignee_user_id": "value",
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "search_match": {
    "field": "title",
    "snippet": "value",
    "highlight_start": 123,
    "highlight_end": 123,
    "turn_index": 123
  }
}`)
	}

	req := receiveWireAgentsTriagetickets(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/triage-tickets/value/turn-comments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/triage-tickets/value/turn-comments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTriagetickets(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAddTurnCommentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTriagetickets(t, 422, `{
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

	var request agents.AddTurnCommentRequestModel
	if err := json.Unmarshal([]byte(`{
  "agentqa_ticket_id": "value",
  "turn_index": 123,
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.AddTurnComment(context.Background(), string("value"), &request)
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

type wireRequestAgentsTriagetickets struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsTriagetickets records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsTriagetickets(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsTriagetickets) {
	t.Helper()
	recorded := make(chan wireRequestAgentsTriagetickets, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsTriagetickets{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsTriagetickets(t *testing.T, recorded <-chan wireRequestAgentsTriagetickets) wireRequestAgentsTriagetickets {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsTriagetickets{}
	}
}

// wireJSONContainsAgentsTriagetickets reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsTriagetickets(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsTriagetickets(got, want)
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

// setWireValueAgentsTriagetickets sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsTriagetickets(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsTriagetickets(got, want []byte) bool {
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
