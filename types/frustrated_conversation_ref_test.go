package types

import (
	"testing"
)

func TestFrustratedConversationRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "agent_id": "value",
  "start_time_unix_secs": 123,
  "overall_label": "positive",
  "overall_sentiment_score": 9.99,
  "overall_frustration_score": 9.99
}`
	assertRoundTrip(t, wire, new(FrustratedConversationRef))
}
