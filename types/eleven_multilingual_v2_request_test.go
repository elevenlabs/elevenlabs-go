package types

import (
	"testing"
)

func TestElevenMultilingualV2RequestRoundTrip(t *testing.T) {
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
  "model_id": "eleven_multilingual_v2",
  "voice_settings": {
    "stability": 0.5,
    "similarity_boost": 0.5,
    "style": 0.5,
    "use_speaker_boost": true,
    "speed": 1.2
  }
}`
	assertRoundTrip(t, wire, new(ElevenMultilingualV2Request))
}
