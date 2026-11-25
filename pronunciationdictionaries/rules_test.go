package pronunciationdictionaries

import (
	"testing"
)

func TestBodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRoundTrip(t *testing.T) {
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
	assertRoundTrip(t, wire, new(BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost))
}

func TestBodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "string_to_replace": "value",
  "case_sensitive": true,
  "word_boundaries": true,
  "type": "alias",
  "alias": "value"
}`
	assertRoundTrip(t, wire, new(BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem))
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
