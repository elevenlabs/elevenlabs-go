package types

import (
	"testing"
)

func TestConversationHistoryEvaluationCriteriaResultCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "criteria_id": "value",
  "result": "success",
  "rationale": "value",
  "scoring_mode": "binary",
  "score": 123,
  "max_score": 123
}`
	assertRoundTrip(t, wire, new(ConversationHistoryEvaluationCriteriaResultCommonModel))
}
