package types

import (
	"testing"
)

func TestBodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segments": [
    "value"
  ],
  "languages": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyDubsAllOrSomeSegmentsAndLanguagesV1DubbingResourceDubbingIdDubPost))
}
