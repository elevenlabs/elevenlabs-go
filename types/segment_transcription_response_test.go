package types

import (
	"testing"
)

func TestSegmentTranscriptionResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentTranscriptionResponse))
}
