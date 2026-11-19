package types

import (
	"testing"
)

func TestRedactedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "redacted_transcript",
  "text": "value",
  "redacted_text": "value"
}`
	assertRoundTrip(t, wire, new(RedactedTranscript))
}
