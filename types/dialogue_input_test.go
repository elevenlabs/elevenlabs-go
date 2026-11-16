package types

import (
	"testing"
)

func TestDialogueInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(DialogueInput))
}
