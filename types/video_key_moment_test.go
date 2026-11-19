package types

import (
	"testing"
)

func TestVideoKeyMomentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "timestamp_ms": 123,
  "type": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(VideoKeyMoment))
}
