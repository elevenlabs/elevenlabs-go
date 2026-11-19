package types

import (
	"testing"
)

func TestPronunciationDictionaryVersionLocatorRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryVersionLocatorRequestModel))
}

func TestPronunciationDictionaryVersionLocatorRequestModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryVersionLocatorRequestModel))
}

func TestPronunciationDictionaryVersionLocatorRequestModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(PronunciationDictionaryVersionLocatorRequestModel))
}
