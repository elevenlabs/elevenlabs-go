package types

import (
	"testing"
)

func TestVideoSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_ms": 123,
  "end_ms": 123,
  "description": "value",
  "subjects": [
    "value"
  ],
  "shot_type": "value",
  "camera_movement": "value",
  "transition_in": "value",
  "has_speech": true,
  "has_music": true,
  "pacing": "value"
}`
	assertRoundTrip(t, wire, new(VideoSegment))
}
