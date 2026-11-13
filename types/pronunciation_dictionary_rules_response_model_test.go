package types

import (
	"testing"
)

func TestPronunciationDictionaryRulesResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "version_id": "value",
  "version_rules_num": 123
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryRulesResponseModel))
}
