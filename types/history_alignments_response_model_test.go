package types

import (
	"testing"
)

func TestHistoryAlignmentsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "alignment": {
    "characters": [
      "value"
    ],
    "character_start_times_seconds": [
      9.99
    ],
    "character_end_times_seconds": [
      9.99
    ]
  },
  "normalized_alignment": {
    "characters": [
      "value"
    ],
    "character_start_times_seconds": [
      9.99
    ],
    "character_end_times_seconds": [
      9.99
    ]
  }
}`
	assertRoundTrip(t, wire, new(HistoryAlignmentsResponseModel))
}
