package types

import (
	"testing"
)

func TestReceiveMessageMultiRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio": "value",
  "is_final": true,
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
  "context_id": "value"
}`
	assertRoundTrip(t, wire, new(ReceiveMessageMulti))
}
