package types

import (
	"testing"
)

func TestDubbingPartialTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "partial_transcript",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(DubbingPartialTranscript))
}
