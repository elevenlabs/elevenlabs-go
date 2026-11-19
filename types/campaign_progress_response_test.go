package types

import (
	"testing"
)

func TestCampaignProgressResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "leads": {
    "total": 123,
    "uncalled": 123,
    "leads_needing_attention": 123,
    "active": 123,
    "finished": 123,
    "finished_by_disposition": {
      "key": 123
    }
  },
  "conversations": {
    "scheduled": 123,
    "scheduled_by_type": {
      "key": 123
    },
    "awaiting_permission": 123,
    "in_progress": 123,
    "awaiting_analysis": 123
  },
  "next_call_time_unix": 123
}`
	assertRoundTrip(t, wire, new(CampaignProgressResponse))
}
