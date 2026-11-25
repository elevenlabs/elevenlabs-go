package types

import (
	"testing"
)

func TestVadScoreVadScoreEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "vad_score": 9.99
}`
	assertRoundTrip(t, wire, new(VadScoreVadScoreEvent))
}
