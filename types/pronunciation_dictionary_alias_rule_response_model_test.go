package types

import (
	"testing"
)

func TestPronunciationDictionaryAliasRuleResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "alias",
  "alias": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryAliasRuleResponseModel))
}
