package types

import (
	"testing"
)

func TestSectionSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "song_id": "value",
  "range": {
    "start_ms": 123,
    "end_ms": 123
  },
  "negative_ranges": [
    {
      "start_ms": 123,
      "end_ms": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(SectionSource))
}
