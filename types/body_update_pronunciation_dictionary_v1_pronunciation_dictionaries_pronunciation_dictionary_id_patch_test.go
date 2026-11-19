package types

import (
	"testing"
)

func TestBodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "archived": true,
  "name": "value"
}`
	assertRoundTrip(t, wire, new(BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch))
}
