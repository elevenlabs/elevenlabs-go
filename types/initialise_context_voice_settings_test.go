package types

import (
	"testing"
)

func TestInitialiseContextVoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 9.99,
  "similarity_boost": 9.99,
  "style": 9.99,
  "use_speaker_boost": true,
  "speed": 9.99
}`
	assertRoundTrip(t, wire, new(InitialiseContextVoiceSettings))
}
