package types

import (
	"testing"
)

func TestAudioResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_event": {
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
  },
  "type": "audio"
}`
	assertRoundTrip(t, wire, new(AudioResponse))
}
