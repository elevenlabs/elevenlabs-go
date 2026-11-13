package types

import (
	"testing"
)

func TestAddPronunciationDictionaryRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "description": "value",
  "workspace_access": "admin"
}`
	assertRoundTrip(t, wire, new(AddPronunciationDictionaryRequest))
}
