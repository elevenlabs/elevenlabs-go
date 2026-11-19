package types

import (
	"testing"
)

func TestAttachedSystemEvaluationRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "system",
  "analysis_item_id": "__system_eval_criteria_sentiment",
  "scope": "conversation",
  "weight": 9.99
}`
	assertRoundTrip(t, wire, new(AttachedSystemEvaluationRef))
}
