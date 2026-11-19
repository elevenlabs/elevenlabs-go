package types

import (
	"testing"
)

func TestLlmParameterEvaluationStrategyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(LlmParameterEvaluationStrategy))
}
