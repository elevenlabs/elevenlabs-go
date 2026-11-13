package types

import (
	"testing"
)

func TestGetTestInvocationsPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "id": "value",
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "runs_diverged_from_version": true,
      "created_at_unix_secs": 123,
      "test_run_count": 123,
      "passed_count": 123,
      "failed_count": 123,
      "pending_count": 123,
      "cancelled_count": 123,
      "title": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "repeat_count": 123,
      "cancelled": true,
      "credits_used": 123,
      "total_price": 9.99
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetTestInvocationsPageResponseModel))
}
