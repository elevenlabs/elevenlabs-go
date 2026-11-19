package types

import (
	"testing"
)

func TestTtsConversationalConfigWorkflowOverrideInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": "eleven_turbo_v2",
  "voice_id": "value",
  "supported_voices": [
    {
      "label": "value",
      "voice_id": "value",
      "description": "value",
      "language": "value",
      "model_family": "turbo",
      "optimize_streaming_latency": 123,
      "stability": 0.5,
      "speed": 1.2,
      "similarity_boost": 0.5,
      "clone_strength": 0.5
    }
  ],
  "expressive_mode": true,
  "suggested_audio_tags": [
    {
      "tag": "value",
      "description": "value"
    }
  ],
  "agent_output_audio_format": "pcm_8000",
  "optimize_streaming_latency": 123,
  "stability": 9.99,
  "speed": 9.99,
  "similarity_boost": 9.99,
  "clone_strength": 9.99,
  "text_normalisation_type": "system_prompt",
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "enable_phoneme_tags": true,
  "audio_effects": {
    "filter_preset_id": "value",
    "distance": 0.5,
    "environment_id": "value",
    "background_noise_id": "value",
    "send_level": 0.5,
    "pan": 0.5
  }
}`
	assertRoundTrip(t, wire, new(TtsConversationalConfigWorkflowOverrideInput))
}
