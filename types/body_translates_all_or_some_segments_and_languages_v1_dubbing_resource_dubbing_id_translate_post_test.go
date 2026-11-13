package types

import (
	"testing"
)

func TestBodyTranslatesAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdTranslatePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segments": [
    "value"
  ],
  "languages": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyTranslatesAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdTranslatePost))
}
