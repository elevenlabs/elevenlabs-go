package types

import (
	"testing"
)

func TestAudioSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_ms": 123,
  "end_ms": 123,
  "description": "value",
  "segment_type": "value",
  "has_speech": true,
  "has_music": true,
  "pacing": "value"
}`
	assertRoundTrip(t, wire, new(AudioSegment))
}
