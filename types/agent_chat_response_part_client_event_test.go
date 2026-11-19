package types

import (
	"testing"
)

func TestAgentChatResponsePartClientEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_response_part": {
    "text": "value",
    "type": "start",
    "event_id": 123,
    "response_id": "value"
  },
  "type": "agent_chat_response_part"
}`
	assertRoundTrip(t, wire, new(AgentChatResponsePartClientEvent))
}
