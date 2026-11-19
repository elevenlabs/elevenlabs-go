package types

import (
	"testing"
)

func TestSystemEvaluationCriteriaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_sentiment_score": 0.5,
  "user_frustration_score": 0.5
}`
	assertRoundTrip(t, wire, new(SystemEvaluationCriteria))
}
