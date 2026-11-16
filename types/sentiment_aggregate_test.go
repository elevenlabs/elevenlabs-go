package types

import (
	"testing"
)

func TestSentimentAggregateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "scored_conversation_count": 123,
  "positive_count": 123,
  "neutral_count": 123,
  "negative_count": 123,
  "average_sentiment_score": 9.99,
  "average_frustration_score": 9.99,
  "recent_scored_conversation_count": 123,
  "recent_positive_count": 123,
  "recent_neutral_count": 123,
  "recent_negative_count": 123,
  "recent_average_sentiment_score": 9.99,
  "recent_average_frustration_score": 9.99
}`
	assertRoundTrip(t, wire, new(SentimentAggregate))
}
