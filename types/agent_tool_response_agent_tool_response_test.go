package types

import (
	"testing"
)

func TestAgentToolResponseAgentToolResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "tool_call_id": "value",
  "tool_type": "value",
  "is_error": true,
  "is_blocked": true,
  "event_id": 123,
  "is_called": true,
  "status": "success"
}`
	assertRoundTrip(t, wire, new(AgentToolResponseAgentToolResponse))
}
