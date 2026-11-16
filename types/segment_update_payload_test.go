package types

import (
	"testing"
)

func TestSegmentUpdatePayloadRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 9.99,
  "end_time": 9.99,
  "text": "value"
}`
	assertRoundTrip(t, wire, new(SegmentUpdatePayload))
}
