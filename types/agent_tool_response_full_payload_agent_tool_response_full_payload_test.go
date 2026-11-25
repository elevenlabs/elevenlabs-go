package types

import (
	"testing"
)

func TestAgentToolResponseFullPayloadAgentToolResponseFullPayloadRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AgentToolResponseFullPayloadAgentToolResponseFullPayload))
}
