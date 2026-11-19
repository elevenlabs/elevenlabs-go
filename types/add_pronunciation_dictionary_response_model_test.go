package types

import (
	"testing"
)

func TestAddPronunciationDictionaryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "created_by": "value",
  "creation_time_unix": 123,
  "version_id": "value",
  "version_rules_num": 123,
  "description": "value",
  "permission_on_resource": "admin"
}`
	assertRoundTrip(t, wire, new(AddPronunciationDictionaryResponseModel))
}
