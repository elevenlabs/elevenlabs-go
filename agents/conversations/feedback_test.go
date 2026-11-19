package conversations

import (
	"testing"
)

func TestConversationFeedbackRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "feedback": "like"
}`
	assertRoundTrip(t, wire, new(ConversationFeedbackRequestModel))
}
