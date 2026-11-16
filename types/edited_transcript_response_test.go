package types

import (
	"testing"
)

func TestEditedTranscriptResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "transcript",
  "text": "value",
  "edited_text": "value"
}`
	assertRoundTrip(t, wire, new(EditedTranscriptResponse))
}
