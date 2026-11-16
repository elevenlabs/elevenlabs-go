package types

import (
	"testing"
)

func TestCompositionPlanChunksItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "duration_ms": 3000,
  "positive_styles": [
    "value"
  ],
  "negative_styles": [
    "value"
  ],
  "context_adherence": "low",
  "conditioning_ref": {
    "song_id": "value",
    "range": {
      "start_ms": 123,
      "end_ms": 123
    }
  },
  "condition_strength": "low"
}`
	assertRoundTrip(t, wire, new(CompositionPlanChunksItem))
}
