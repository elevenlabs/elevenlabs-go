package types

import (
	"testing"
)

func TestVideoAnalysisRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "processing",
  "data": {
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
  },
  "updated_at_ms": 123
}`
	assertRoundTrip(t, wire, new(VideoAnalysis))
}
