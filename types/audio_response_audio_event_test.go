package types

import (
	"testing"
)

func TestAudioResponseAudioEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_base_64": "value",
  "event_id": 123,
  "alignment": {
    "chars": [
      "value"
    ],
    "char_start_times_ms": [
      123
    ],
    "char_durations_ms": [
      123
    ]
  },
  "is_final": true
}`
	assertRoundTrip(t, wire, new(AudioResponseAudioEvent))
}
