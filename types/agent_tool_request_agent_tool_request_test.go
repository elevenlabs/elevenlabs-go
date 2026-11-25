package types

import (
	"testing"
)

func TestAgentToolRequestAgentToolRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "tool_call_id": "value",
  "tool_type": "value",
  "event_id": 123,
  "expects_response": true,
  "disable_interruptions": true,
  "response_timeout_secs": 123,
  "execution_mode": "value"
}`
	assertRoundTrip(t, wire, new(AgentToolRequestAgentToolRequest))
}
