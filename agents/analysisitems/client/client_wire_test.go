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

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "analysis_items": [
    {
      "id": "value",
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "current_version_id": "value",
      "legacy_source_id": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	searchValue := string("value")
	pageSizeValue := int(100)
	createdByUserIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListAnalysisItemsRequest{
		Search:          &searchValue,
		PageSize:        &pageSizeValue,
		CreatedByUserId: &createdByUserIdValue,
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
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "analysis_items": [
    {
      "id": "value",
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "current_version_id": "value",
      "legacy_source_id": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "analysis_items": [
    {
      "id": "value",
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "current_version_id": "value",
      "legacy_source_id": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/analysis-items" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("created_by_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "created_by_user_id", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	searchValue := string("value")
	pageSizeValue := int(100)
	createdByUserIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListAnalysisItemsRequest{
		Search:          &searchValue,
		PageSize:        &pageSizeValue,
		CreatedByUserId: &createdByUserIdValue,
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

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateAnalysisItemsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
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
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/analysis-items" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsAnalysisitemsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	var request agents.CreateAnalysisItemsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
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

func TestUpdateAttachmentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "results": [
    {
      "status": "success",
      "data": {
        "agent_id": "value",
        "branch_id": "value",
        "action": "attach",
        "changed": true
      }
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdateAnalysisItemAttachmentsRequestModel
	if err := json.Unmarshal([]byte(`{
  "analysis_item_id": "value",
  "source": "user",
  "attach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "detach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateAttachments(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "results": [
    {
      "status": "success",
      "data": {
        "agent_id": "value",
        "branch_id": "value",
        "action": "attach",
        "changed": true
      }
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "results": [
    {
      "status": "success",
      "data": {
        "agent_id": "value",
        "branch_id": "value",
        "action": "attach",
        "changed": true
      }
    }
  ]
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/analysis-items/value/attachments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value/attachments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsAnalysisitemsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateAttachmentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	var request agents.UpdateAnalysisItemAttachmentsRequestModel
	if err := json.Unmarshal([]byte(`{
  "analysis_item_id": "value",
  "source": "user",
  "attach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "detach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateAttachments(context.Background(), string("value"), &request)
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

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
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
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/analysis-items/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdateAnalysisItemsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
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
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/analysis-items/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsAnalysisitemsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	var request agents.UpdateAnalysisItemsRequestBody
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
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

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	forceValue := bool(true)
	request := &agents.DeleteAnalysisItemsRequest{
		Force: &forceValue,
	}
	result, err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/analysis-items/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value")
	}
	if got := req.Query.Get("force"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "force", got, "true")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	forceValue := bool(true)
	request := &agents.DeleteAnalysisItemsRequest{
		Force: &forceValue,
	}
	_, err := client.Delete(context.Background(), string("value"), request)
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

func TestListVersionsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "versions": [
    {
      "id": "value",
      "version_number": 123,
      "created_at_unix_secs": 123,
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      }
    }
  ],
  "current_version_id": "value",
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.ListVersionsAnalysisItemsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.ListVersions(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "versions": [
    {
      "id": "value",
      "version_number": 123,
      "created_at_unix_secs": 123,
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      }
    }
  ],
  "current_version_id": "value",
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "versions": [
    {
      "id": "value",
      "version_number": 123,
      "created_at_unix_secs": 123,
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      }
    }
  ],
  "current_version_id": "value",
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/analysis-items/value/versions" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value/versions")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListVersionsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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
	cursorValue := string("value")
	request := &agents.ListVersionsAnalysisItemsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	_, err := client.ListVersions(context.Background(), string("value"), request)
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

func TestCopyWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsAnalysisitemsClient(t, 200, `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Copy(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsAnalysisitemsClient(gotResponse, []byte(`{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "definition": {
    "name": "value",
    "source": "user",
    "llm": "gpt-4o-mini",
    "kind": "evaluation",
    "conversation_goal_prompt": "value",
    "use_knowledge_base": true,
    "scoring": {
      "mode": "binary"
    }
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`)
	}

	req := receiveWireAgentsAnalysisitemsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/analysis-items/value/copy" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/analysis-items/value/copy")
	}
}

func TestCopyWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsAnalysisitemsClient(t, 422, `{
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

	_, err := client.Copy(context.Background(), string("value"))
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

type wireRequestAgentsAnalysisitemsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsAnalysisitemsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsAnalysisitemsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsAnalysisitemsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsAnalysisitemsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsAnalysisitemsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsAnalysisitemsClient(t *testing.T, recorded <-chan wireRequestAgentsAnalysisitemsClient) wireRequestAgentsAnalysisitemsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsAnalysisitemsClient{}
	}
}

// wireJSONContainsAgentsAnalysisitemsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsAnalysisitemsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsAnalysisitemsClient(got, want)
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

// setWireValueAgentsAnalysisitemsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsAnalysisitemsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsAnalysisitemsClient(got, want []byte) bool {
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
