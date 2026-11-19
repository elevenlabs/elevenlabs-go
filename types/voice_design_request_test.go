package types

import (
	"testing"
)

func TestVoiceDesignRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "quality": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "should_enhance": true
}`
	assertRoundTrip(t, wire, new(VoiceDesignRequest))
}
