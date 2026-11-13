package types

import (
	"testing"
)

func TestElevenFlashV25RequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "text": "value",
  "voice": "value",
  "output_format": "mp3_22050_32",
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "model_id": "eleven_flash_v2_5",
  "language_code": "value",
  "voice_settings": {
    "stability": 0.5,
    "similarity_boost": 0.5,
    "speed": 1.2
  }
}`
	assertRoundTrip(t, wire, new(ElevenFlashV25Request))
}
