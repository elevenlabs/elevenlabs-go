package types

import (
	"testing"
)

func TestInitialiseContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "voice_settings": {
    "stability": 9.99,
    "similarity_boost": 9.99,
    "style": 9.99,
    "use_speaker_boost": true,
    "speed": 9.99
  },
  "generation_config": {
    "chunk_length_schedule": [
      123
    ]
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "xi_api_key": "value",
  "authorization": "value",
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(InitialiseContext))
}
