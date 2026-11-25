package types

import (
	"testing"
)

func TestAgentToolResponseFullPayloadRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_tool_response_full_payload": {
    "tool_name": "value",
    "tool_call_id": "value",
    "tool_type": "value",
    "is_error": true,
    "is_blocked": true,
    "event_id": 123,
    "is_called": true,
    "full_tool_result": "value",
    "truncated": true,
    "status": "success"
  },
  "type": "agent_tool_response_full_payload"
}`
	assertRoundTrip(t, wire, new(AgentToolResponseFullPayload))
}
