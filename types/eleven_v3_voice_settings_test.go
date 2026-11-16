package types

import (
	"testing"
)

func TestElevenV3VoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5
}`
	assertRoundTrip(t, wire, new(ElevenV3VoiceSettings))
}
