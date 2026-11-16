package types

import (
	"testing"
)

func TestVideoAnalysisResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "content_type": "value",
  "overall_pacing": "value",
  "subjects": [
    {
      "name": "value",
      "description": "value"
    }
  ],
  "segments": [
    {
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
    }
  ],
  "key_moments": [
    {
      "timestamp_ms": 123,
      "type": "value",
      "description": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(VideoAnalysisResult))
}
