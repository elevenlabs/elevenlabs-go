package types

import (
	"testing"
)

func TestPronunciationDictionaryVersionResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version_id": "value",
  "version_rules_num": 123,
  "pronunciation_dictionary_id": "value",
  "dictionary_name": "value",
  "version_name": "value",
  "permission_on_resource": "admin",
  "created_by": "value",
  "creation_time_unix": 123,
  "archived_time_unix": 123
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryVersionResponseModel))
}
