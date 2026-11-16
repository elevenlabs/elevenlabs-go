package types

import (
	"testing"
)

func TestRealtimeConversationGroupResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(RealtimeConversationGroupResponseModel))
}
