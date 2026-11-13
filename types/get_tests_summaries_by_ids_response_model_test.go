package types

import (
	"testing"
)

func TestGetTestsSummariesByIdsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tests": {
    "key": {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  }
}`
	assertRoundTrip(t, wire, new(GetTestsSummariesByIdsResponseModel))
}
