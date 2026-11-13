package types

import (
	"testing"
)

func TestTextToDialogueMultiClientMessageVoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5
}`
	assertRoundTrip(t, wire, new(TextToDialogueMultiClientMessageVoiceSettings))
}
