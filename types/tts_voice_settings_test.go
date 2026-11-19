package types

import (
	"testing"
)

func TestTtsVoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "similarity_boost": 0.5,
  "style": 0.5,
  "use_speaker_boost": true,
  "speed": 1.2
}`
	assertRoundTrip(t, wire, new(TtsVoiceSettings))
}
