package types

import (
	"testing"
)

func TestGetAgentsPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agents": [
    {
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
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetAgentsPageResponseModel))
}
