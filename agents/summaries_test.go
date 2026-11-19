package agents

import (
	"testing"
)

func TestGetSummariesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "success",
  "data": {
    "agent_id": "value",
    "name": "value",
    "voice_id": "value",
    "tags": [
      "value"
    ],
    "created_at_unix_secs": 123,
    "updated_at_unix_secs": 123,
    "access_info": {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    },
    "last_call_time_unix_secs": 123,
    "archived": true
  }
}`
	assertRoundTrip(t, wire, new(GetSummariesResponse))
}
