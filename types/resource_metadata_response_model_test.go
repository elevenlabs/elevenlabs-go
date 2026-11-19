package types

import (
	"testing"
)

func TestResourceMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resource_id": "value",
  "resource_name": "value",
  "resource_type": "voice",
  "creator_user_id": "value",
  "anonymous_access_level_override": "admin",
  "role_to_group_ids": {
    "key": [
      "value"
    ]
  },
  "share_options": [
    {
      "name": "value",
      "id": "value",
      "type": "user"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ResourceMetadataResponseModel))
}
