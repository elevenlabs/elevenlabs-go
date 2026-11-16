package types

import (
	"testing"
)

func TestRegexParameterEvaluationStrategyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "regex",
  "pattern": "value"
}`
	assertRoundTrip(t, wire, new(RegexParameterEvaluationStrategy))
}
