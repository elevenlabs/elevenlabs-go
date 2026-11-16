package types

import (
	"testing"
)

func TestConversationHistoryFeedbackCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "thumbs",
  "overall_score": "like",
  "likes": 123,
  "dislikes": 123,
  "rating": 123,
  "comment": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryFeedbackCommonModel))
}
