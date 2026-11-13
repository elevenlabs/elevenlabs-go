package types

import (
	"testing"
)

func TestMatchAnythingParameterEvaluationStrategyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "anything"
}`
	assertRoundTrip(t, wire, new(MatchAnythingParameterEvaluationStrategy))
}
