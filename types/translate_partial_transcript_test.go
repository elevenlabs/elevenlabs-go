package types

import (
	"testing"
)

func TestTranslatePartialTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "partial_transcript",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(TranslatePartialTranscript))
}
