package types

import (
	"testing"
)

func TestAgentBranchSummaryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "agent_id": "value",
  "description": "value",
  "created_at": 123,
  "last_committed_at": 123,
  "is_archived": true,
  "protection_status": "writer_perms_required",
  "auto_rebase": true,
  "sync_status": {
    "last_attempt": {
      "outcome": "succeeded",
      "attempted_at": 123
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
  "current_live_percentage": 9.99,
  "parent_branch_id": "value",
  "draft_exists": true,
  "draft_created_at": 123,
  "draft_is_behind_tip": true,
  "calls_7d": 123,
  "commits_ahead": 123,
  "commits_behind": 123,
  "merged_into_branch_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentBranchSummary))
}
