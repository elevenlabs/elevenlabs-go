package mergeproposals

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

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "results": [
    {
      "id": "value",
      "agent_id": "value",
      "source_branch_id": "value",
      "target_branch_id": "value",
      "source_tip_version_id_at_creation": "value",
      "title": "value",
      "description": "value",
      "requested_reviewer_user_ids": [
        "value"
      ],
      "reviews": [
        {
          "user_id": "value",
          "state": "approved",
          "submitted_at": 123,
          "comment": "value",
          "reviewer_role": "admin"
        }
      ],
      "comments": [
        {
          "id": "value",
          "user_id": "value",
          "body": "value",
          "created_at": 123
        }
      ],
      "outcome": {
        "status": "open"
      },
      "created_at": 123,
      "updated_at": 123,
      "author_user_id": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	sourceBranchIdValue := string("value")
	targetBranchIdValue := string("value")
	searchValue := string("value")
	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.ListMergeProposalsRequest{
		SourceBranchId: &sourceBranchIdValue,
		TargetBranchId: &targetBranchIdValue,
		Search:         &searchValue,
		PageSize:       &pageSizeValue,
		Cursor:         &cursorValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "results": [
    {
      "id": "value",
      "agent_id": "value",
      "source_branch_id": "value",
      "target_branch_id": "value",
      "source_tip_version_id_at_creation": "value",
      "title": "value",
      "description": "value",
      "requested_reviewer_user_ids": [
        "value"
      ],
      "reviews": [
        {
          "user_id": "value",
          "state": "approved",
          "submitted_at": 123,
          "comment": "value",
          "reviewer_role": "admin"
        }
      ],
      "comments": [
        {
          "id": "value",
          "user_id": "value",
          "body": "value",
          "created_at": 123
        }
      ],
      "outcome": {
        "status": "open"
      },
      "created_at": 123,
      "updated_at": 123,
      "author_user_id": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "results": [
    {
      "id": "value",
      "agent_id": "value",
      "source_branch_id": "value",
      "target_branch_id": "value",
      "source_tip_version_id_at_creation": "value",
      "title": "value",
      "description": "value",
      "requested_reviewer_user_ids": [
        "value"
      ],
      "reviews": [
        {
          "user_id": "value",
          "state": "approved",
          "submitted_at": 123,
          "comment": "value",
          "reviewer_role": "admin"
        }
      ],
      "comments": [
        {
          "id": "value",
          "user_id": "value",
          "body": "value",
          "created_at": 123
        }
      ],
      "outcome": {
        "status": "open"
      },
      "created_at": 123,
      "updated_at": 123,
      "author_user_id": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals")
	}
	if got := req.Query.Get("source_branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "source_branch_id", got, "value")
	}
	if got := req.Query.Get("target_branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "target_branch_id", got, "value")
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

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	sourceBranchIdValue := string("value")
	targetBranchIdValue := string("value")
	searchValue := string("value")
	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.ListMergeProposalsRequest{
		SourceBranchId: &sourceBranchIdValue,
		TargetBranchId: &targetBranchIdValue,
		Search:         &searchValue,
		PageSize:       &pageSizeValue,
		Cursor:         &cursorValue,
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "created_merge_proposal_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "triage_ticket_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "created_merge_proposal_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "created_merge_proposal_id": "value"
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMergeproposals(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	var request agents.BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost
	if err := json.Unmarshal([]byte(`{
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "triage_ticket_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), string("value"), &request)
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

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"), string("value2"))
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

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "close": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals/value2")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMergeproposals(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	var request agents.BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "close": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), string("value2"), &request)
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

func TestSubmitReviewWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "state": "approved",
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.SubmitReview(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals/value2/reviews" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals/value2/reviews")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMergeproposals(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSubmitReviewWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	var request agents.BodyReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "state": "approved",
  "comment": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.SubmitReview(context.Background(), string("value"), string("value2"), &request)
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

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyCommentOnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "body": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.AddComment(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals/value2/comments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals/value2/comments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMergeproposals(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAddCommentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	var request agents.BodyCommentOnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "body": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.AddComment(context.Background(), string("value"), string("value2"), &request)
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

	server, recorded := serveWireAgentsMergeproposals(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "archive_source_branch": true,
  "force": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Merge(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMergeproposals(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "source_branch_id": "value",
  "target_branch_id": "value",
  "source_tip_version_id_at_creation": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "reviews": [
    {
      "user_id": "value",
      "state": "approved",
      "submitted_at": 123,
      "comment": "value",
      "reviewer_role": "admin"
    }
  ],
  "comments": [
    {
      "id": "value",
      "user_id": "value",
      "body": "value",
      "created_at": 123
    }
  ],
  "outcome": {
    "status": "open"
  },
  "created_at": 123,
  "updated_at": 123,
  "author_user_id": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`)
	}

	req := receiveWireAgentsMergeproposals(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agents/value/merge-proposals/value2/merge" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/merge-proposals/value2/merge")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMergeproposals(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestMergeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMergeproposals(t, 422, `{
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

	var request agents.BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost
	if err := json.Unmarshal([]byte(`{
  "merge_proposal_id": "value",
  "agent_id": "value",
  "archive_source_branch": true,
  "force": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Merge(context.Background(), string("value"), string("value2"), &request)
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

type wireRequestAgentsMergeproposals struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsMergeproposals records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsMergeproposals(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsMergeproposals) {
	t.Helper()
	recorded := make(chan wireRequestAgentsMergeproposals, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsMergeproposals{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsMergeproposals(t *testing.T, recorded <-chan wireRequestAgentsMergeproposals) wireRequestAgentsMergeproposals {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsMergeproposals{}
	}
}

// wireJSONContainsAgentsMergeproposals reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsMergeproposals(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsMergeproposals(got, want)
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

// setWireValueAgentsMergeproposals sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsMergeproposals(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsMergeproposals(got, want []byte) bool {
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
