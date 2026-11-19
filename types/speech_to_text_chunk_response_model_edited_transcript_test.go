package types

import (
	"testing"
)

func TestSpeechToTextChunkResponseModelEditedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "transcript",
  "text": "value",
  "edited_text": "value"
}`
	assertRoundTrip(t, wire, new(SpeechToTextChunkResponseModelEditedTranscript))
}
