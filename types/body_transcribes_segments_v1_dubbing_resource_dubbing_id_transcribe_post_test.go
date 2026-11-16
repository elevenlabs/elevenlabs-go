package types

import (
	"testing"
)

func TestBodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segments": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyTranscribesSegmentsV1DubbingResourceDubbingIdTranscribePost))
}
