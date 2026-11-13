package types

import (
	"testing"
)

func TestPaginatedResultAgentMergeProposalResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(PaginatedResultAgentMergeProposalResponse))
}
