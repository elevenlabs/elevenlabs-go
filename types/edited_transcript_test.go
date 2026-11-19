package types

import (
	"testing"
)

func TestEditedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "value",
  "text": "value",
  "edited_text": "value"
}`
	assertRoundTrip(t, wire, new(EditedTranscript))
}
