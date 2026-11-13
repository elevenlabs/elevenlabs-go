package types

import (
	"testing"
)

func TestBodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "inputs": [
    {
      "text": "value",
      "voice_id": "value"
    }
  ],
  "model_id": "value",
  "language_code": "value",
  "settings": {
    "stability": 0.5,
    "similarity": 0.5,
    "clone_strength": 0.5
  },
  "previous_text": "value",
  "future_text": "value",
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "apply_text_normalization": "auto",
  "use_pvc_as_ivc": true,
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "audio_effects": {
    "filter_preset_id": "value",
    "distance": 0.5,
    "environment_id": "value",
    "background_noise_id": "value",
    "send_level": 0.5,
    "pan": 0.5
  }
}`
	assertRoundTrip(t, wire, new(BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPost))
}
