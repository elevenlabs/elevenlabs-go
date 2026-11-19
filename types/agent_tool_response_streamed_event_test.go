package types

import (
	"testing"
)

func TestAgentToolResponseStreamedEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_tool_response": {
    "tool_name": "value",
    "tool_call_id": "value",
    "tool_type": "value",
    "is_error": true,
    "is_blocked": true,
    "event_id": 123,
    "is_called": true,
    "status": "success"
  },
  "type": "agent_tool_response"
}`
	assertRoundTrip(t, wire, new(AgentToolResponseStreamedEvent))
}
