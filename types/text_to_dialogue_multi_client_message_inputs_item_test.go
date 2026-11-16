package types

import (
	"testing"
)

func TestTextToDialogueMultiClientMessageInputsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "voice_id": "value",
  "new_turn": true
}`
	assertRoundTrip(t, wire, new(TextToDialogueMultiClientMessageInputsItem))
}
