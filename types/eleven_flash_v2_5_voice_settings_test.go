package types

import (
	"testing"
)

func TestElevenFlashV25VoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "similarity_boost": 0.5,
  "speed": 1.2
}`
	assertRoundTrip(t, wire, new(ElevenFlashV25VoiceSettings))
}
