package types

import (
	"testing"
)

func TestDefaultSharingGroupResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "group": {
    "name": "value",
    "id": "value",
    "members": [
      "value"
    ],
    "permissions": [
      "text_to_speech"
    ],
    "group_usage_limit": 123,
    "group_pvc_limit": 123,
    "group_concurrency_limit": 123,
    "group_dubbing_concurrency_limit": 123,
    "group_music_concurrency_limit": 123,
    "character_count": 123,
    "is_scim_synced": true,
    "scim_group": {
      "scim_external_id": "value",
      "display_name": "value",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "seat_type": "workspace_admin"
    },
    "scim_frozen": true
  },
  "permission_level": "admin"
}`
	assertRoundTrip(t, wire, new(DefaultSharingGroupResponseModel))
}
