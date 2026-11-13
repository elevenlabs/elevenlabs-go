package types

import (
	"testing"
)

func TestTranslateFinalTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "final_transcript",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(TranslateFinalTranscript))
}
