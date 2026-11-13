package types

import (
	"testing"
)

func TestVoiceSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`
	assertRoundTrip(t, wire, new(VoiceSettingsResponseModel))
}

func TestVoiceSettingsResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`
	assertRoundTrip(t, wire, new(VoiceSettingsResponseModel))
}

func TestVoiceSettingsResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`
	assertRoundTrip(t, wire, new(VoiceSettingsResponseModel))
}
