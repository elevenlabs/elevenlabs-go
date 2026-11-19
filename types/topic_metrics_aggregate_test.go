package types

import (
	"testing"
)

func TestTopicMetricsAggregateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_count": 123,
  "sentiment": {
    "sentiment": {
      "count": 123,
      "sum": 9.99,
      "min": 9.99,
      "max": 9.99
    },
    "frustration": {
      "count": 123,
      "sum": 9.99,
      "min": 9.99,
      "max": 9.99
    },
    "positive_count": 123,
    "neutral_count": 123,
    "negative_count": 123
  },
  "evaluation_criteria": [
    {
      "criteria_id": "value",
      "success_count": 123,
      "failure_count": 123,
      "unknown_count": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(TopicMetricsAggregate))
}
