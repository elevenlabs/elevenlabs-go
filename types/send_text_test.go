package types

import (
	"testing"
)

func TestSendTextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "try_trigger_generation": true,
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
  "flush": true
}`
	assertRoundTrip(t, wire, new(SendText))
}
