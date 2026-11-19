package projects

import (
	"testing"
)

func TestBodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "invalidate_affected_text": true
}`
	assertRoundTrip(t, wire, new(BodyCreatePronunciationDictionariesV1StudioProjectsProjectIdPronunciationDictionariesPost))
}
