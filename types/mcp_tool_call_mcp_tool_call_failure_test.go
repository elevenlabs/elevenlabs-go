package types

import (
	"testing"
)

func TestMcpToolCallMcpToolCallFailureRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service_id": "value",
  "tool_call_id": "value",
  "tool_name": "value",
  "tool_description": "value",
  "parameters": {},
  "timestamp": "value",
  "state": "failure",
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(McpToolCallMcpToolCallFailure))
}
