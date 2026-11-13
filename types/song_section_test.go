package types

import (
	"testing"
)

func TestSongSectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "section_name": "value",
  "positive_local_styles": [
    "value"
  ],
  "negative_local_styles": [
    "value"
  ],
  "duration_ms": 3000,
  "lines": [
    "value"
  ],
  "source_from": {
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
  }
}`
	assertRoundTrip(t, wire, new(SongSection))
}
