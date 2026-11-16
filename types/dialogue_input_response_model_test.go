package types

import (
	"testing"
)

func TestDialogueInputResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "voice_id": "value",
  "voice_name": "value"
}`
	assertRoundTrip(t, wire, new(DialogueInputResponseModel))
}
