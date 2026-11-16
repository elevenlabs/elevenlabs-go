package types

import (
	"testing"
)

func TestGetPronunciationDictionariesMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionaries": [
    {
      "id": "value",
      "latest_version_id": "value",
      "latest_version_rules_num": 123,
      "name": "value",
      "permission_on_resource": "admin",
      "created_by": "value",
      "creation_time_unix": 123,
      "archived_time_unix": 123,
      "description": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetPronunciationDictionariesMetadataResponseModel))
}
