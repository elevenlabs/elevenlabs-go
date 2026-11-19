package types

import (
	"testing"
)

func TestConversationTurnRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "turn_id": "value",
  "status": "in_progress",
  "output": [
    {
      "type": "text",
      "role": "agent",
      "id": "value",
      "text": "value",
      "created_at_unix_ms": 123
    }
  ],
  "error": {
    "type": "value",
    "code": "value",
    "message": "value"
  },
  "required_actions": [
    {
      "type": "client_tool_call",
      "tool_call_id": "value",
      "tool_name": "value",
      "parameters": {
        "key": {}
      }
    }
  ],
  "stream_cursor": "value"
}`
	assertRoundTrip(t, wire, new(ConversationTurn))
}
