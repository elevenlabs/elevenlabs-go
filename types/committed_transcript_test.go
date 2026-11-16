package types

import (
	"testing"
)

func TestCommittedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "committed_transcript",
  "text": "value",
  "segment_index": 123
}`
	assertRoundTrip(t, wire, new(CommittedTranscript))
}
