package types

import (
	"testing"
)

func TestVadScoreRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "vad_score_event": {
    "vad_score": 9.99
  },
  "type": "vad_score"
}`
	assertRoundTrip(t, wire, new(VadScore))
}
