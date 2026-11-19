package types

import (
	"testing"
)

func TestDubbingTranscriptCharacterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptCharacter))
}
