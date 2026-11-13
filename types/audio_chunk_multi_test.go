package types

import (
	"testing"
)

func TestAudioChunkMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio": "value",
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
  "normalized_alignment": {
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
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(AudioChunkMulti))
}
