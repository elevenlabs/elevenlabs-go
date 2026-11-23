package types

import (
	"testing"
)

func TestConversationAgentResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_response_event": {
    "agent_response": "value",
    "event_id": 123,
    "response_id": "value",
    "in_response_to_ids": [
      "value"
    ],
    "response_type": "text",
    "attachments": [
      {
        "url": "value",
        "name": "value",
        "mime_type": "value"
      }
    ]
  },
  "type": "agent_response"
}`
	assertRoundTrip(t, wire, new(ConversationAgentResponse))
}
