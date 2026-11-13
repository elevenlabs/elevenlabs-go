package types

import (
	"testing"
)

func TestBodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "rules": [
    {
      "string_to_replace": "value",
      "case_sensitive": true,
      "word_boundaries": true,
      "type": "alias",
      "alias": "value"
    }
  ],
  "name": "value",
  "description": "value",
  "workspace_access": "admin"
}`
	assertRoundTrip(t, wire, new(BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost))
}
