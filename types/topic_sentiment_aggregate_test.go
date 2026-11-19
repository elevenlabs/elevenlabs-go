package types

import (
	"testing"
)

func TestTopicSentimentAggregateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(TopicSentimentAggregate))
}
