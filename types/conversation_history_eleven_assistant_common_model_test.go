package types

import (
	"testing"
)

func TestConversationHistoryElevenAssistantCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_eleven_assistant": true
}`
	assertRoundTrip(t, wire, new(ConversationHistoryElevenAssistantCommonModel))
}
