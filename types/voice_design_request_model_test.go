package types

import (
	"testing"
)

func TestVoiceDesignRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "model_id": "eleven_multilingual_ttv_v2",
  "text": "valuexxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "auto_generate_text": true,
  "loudness": 0.5,
  "seed": 123,
  "guidance_scale": 9.99,
  "stream_previews": true,
  "should_enhance": true,
  "remixing_session_id": "value",
  "remixing_session_iteration_id": "value",
  "quality": 0.5,
  "reference_audio_base64": "value",
  "prompt_strength": 0.5
}`
	assertRoundTrip(t, wire, new(VoiceDesignRequestModel))
}
