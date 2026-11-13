package pronunciationdictionaries

import (
	"testing"
)

func TestPronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRoundTrip(t *testing.T) {
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
  ]
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost))
}

func TestDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "alias",
  "alias": "value"
}`
	assertRoundTrip(t, wire, new(DictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem))
}

func TestPronunciationDictionaryRoundTrip(t *testing.T) {
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
  ]
}`
	assertRoundTrip(t, wire, new(PronunciationDictionary))
}

func TestPronunciationDictionaryRuleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "alias",
  "alias": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryRule))
}

func TestRemovePronunciationDictionaryRulesRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "rule_strings": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(RemovePronunciationDictionaryRulesRequest))
}
