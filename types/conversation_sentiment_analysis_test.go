package types

import (
	"testing"
)

func TestConversationSentimentAnalysisRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "overall_label": "positive",
  "overall_sentiment_score": 0.5,
  "overall_frustration_score": 0.5,
  "min_user_sentiment_score": 0.5,
  "max_user_frustration_score": 0.5,
  "num_scored_user_turns": 123
}`
	assertRoundTrip(t, wire, new(ConversationSentimentAnalysis))
}
