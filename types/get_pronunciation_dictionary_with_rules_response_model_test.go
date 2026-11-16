package types

import (
	"testing"
)

func TestGetPronunciationDictionaryWithRulesResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "latest_version_id": "value",
  "latest_version_rules_num": 123,
  "name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123,
  "description": "value",
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetPronunciationDictionaryWithRulesResponseModel))
}
