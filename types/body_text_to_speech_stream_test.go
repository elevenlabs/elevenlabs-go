package types

import (
	"testing"
)

func TestBodyTextToSpeechStreamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "model_id": "value",
  "language_code": "value",
  "voice_settings": {
    "stability": 0.5,
    "use_speaker_boost": true,
    "similarity_boost": 0.5,
    "style": 9.99,
    "speed": 9.99,
    "clone_strength": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "previous_text": "value",
  "next_text": "value",
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "use_pvc_as_ivc": true,
  "apply_text_normalization": "auto",
  "apply_language_text_normalization": true,
  "audio_effects": {
    "filter_preset_id": "value",
    "distance": 0.5,
    "environment_id": "value",
    "background_noise_id": "value",
    "send_level": 0.5,
    "pan": 0.5
  }
}`
	assertRoundTrip(t, wire, new(BodyTextToSpeechStream))
}
