package types

import (
	"testing"
)

func TestBodyComposeMusicV1MusicPostCompositionPlanRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "positive_global_styles": [
    "value"
  ],
  "negative_global_styles": [
    "value"
  ],
  "sections": [
    {
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
    }
  ]
}`
	assertRoundTrip(t, wire, new(BodyComposeMusicV1MusicPostCompositionPlan))
}
