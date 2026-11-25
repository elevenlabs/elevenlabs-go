package types

import (
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideConversationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": 123
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataConversationConfigOverrideConversation))
}
