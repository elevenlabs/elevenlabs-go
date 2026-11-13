package types

import (
	"testing"
)

func TestSegmentSubtitleFrameRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 9.99,
  "end_time": 9.99,
  "lines": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SegmentSubtitleFrame))
}
