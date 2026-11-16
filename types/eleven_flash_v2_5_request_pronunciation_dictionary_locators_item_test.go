package types

import (
	"testing"
)

func TestElevenFlashV25RequestPronunciationDictionaryLocatorsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(ElevenFlashV25RequestPronunciationDictionaryLocatorsItem))
}
