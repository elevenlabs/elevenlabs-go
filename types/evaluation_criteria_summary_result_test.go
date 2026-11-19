package types

import (
	"testing"
)

func TestEvaluationCriteriaSummaryResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result": "success",
  "score": 123,
  "max_score": 123
}`
	assertRoundTrip(t, wire, new(EvaluationCriteriaSummaryResult))
}
