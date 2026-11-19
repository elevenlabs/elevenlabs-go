package types

import (
	"testing"
)

func TestAudioKeyMomentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "timestamp_ms": 123,
  "type": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(AudioKeyMoment))
}
