package types

import (
	"testing"
)

func TestAsrConversationalConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "quality": "high",
  "provider": "elevenlabs",
  "user_input_audio_format": "pcm_8000",
  "keywords": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(AsrConversationalConfig))
}

func TestAsrConversationalConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "quality": "high",
  "provider": "elevenlabs",
  "user_input_audio_format": "pcm_8000",
  "keywords": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(AsrConversationalConfig))
}
