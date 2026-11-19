package types

import (
	"testing"
)

func TestVoiceRemixRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_description": "value",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "prompt_strength": 0.5
}`
	assertRoundTrip(t, wire, new(VoiceRemixRequestModel))
}
