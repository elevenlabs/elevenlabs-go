package types

import (
	"testing"
)

func TestExactParameterEvaluationStrategyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "exact",
  "expected_value": "value"
}`
	assertRoundTrip(t, wire, new(ExactParameterEvaluationStrategy))
}
