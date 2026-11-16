package types

import (
	"testing"
)

func TestNumericUniformScoringRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "numeric_uniform",
  "max_score": 123,
  "instructions": "value"
}`
	assertRoundTrip(t, wire, new(NumericUniformScoring))
}
