package types

import (
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "name": "value",
  "id": "value",
  "usage_mode": "prompt"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem))
}
