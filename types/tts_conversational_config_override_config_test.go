package types

import (
	"testing"
)

func TestTtsConversationalConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": true,
  "voice_id": true,
  "supported_voices": true,
  "stability": true,
  "speed": true,
  "similarity_boost": true,
  "clone_strength": true,
  "pronunciation_dictionary_locators": true
}`
	assertRoundTrip(t, wire, new(TtsConversationalConfigOverrideConfig))
}
