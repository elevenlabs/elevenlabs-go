package types

import (
	"testing"
)

func TestPronunciationDictionaryPhonemeRuleRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "phoneme",
  "phoneme": "value",
  "alphabet": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryPhonemeRuleRequestModel))
}

func TestPronunciationDictionaryPhonemeRuleRequestModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "phoneme",
  "phoneme": "value",
  "alphabet": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryPhonemeRuleRequestModel))
}
