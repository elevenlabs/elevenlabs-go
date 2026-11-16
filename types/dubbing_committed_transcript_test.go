package types

import (
	"testing"
)

func TestDubbingCommittedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "committed_transcript",
  "text": "value",
  "segment_id": "value",
  "start_s": 9.99,
  "end_s": 9.99,
  "speaker_id": "value"
}`
	assertRoundTrip(t, wire, new(DubbingCommittedTranscript))
}
