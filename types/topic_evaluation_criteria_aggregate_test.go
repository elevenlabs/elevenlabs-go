package types

import (
	"testing"
)

func TestTopicEvaluationCriteriaAggregateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "criteria_id": "value",
  "success_count": 123,
  "failure_count": 123,
  "unknown_count": 123
}`
	assertRoundTrip(t, wire, new(TopicEvaluationCriteriaAggregate))
}
