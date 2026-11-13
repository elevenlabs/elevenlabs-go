package types

import (
	"testing"
)

func TestSegmentCreatePayloadRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 9.99,
  "end_time": 9.99,
  "text": "value",
  "translations": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(SegmentCreatePayload))
}
