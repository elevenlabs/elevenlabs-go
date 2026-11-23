package types

import (
	"testing"
)

func TestAgentToolRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_tool_request": {
    "tool_name": "value",
    "tool_call_id": "value",
    "tool_type": "value",
    "event_id": 123,
    "expects_response": true,
    "disable_interruptions": true,
    "response_timeout_secs": 123,
    "execution_mode": "value"
  },
  "type": "agent_tool_request"
}`
	assertRoundTrip(t, wire, new(AgentToolRequest))
}
