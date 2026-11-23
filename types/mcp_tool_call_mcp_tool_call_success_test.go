package types

import (
	"testing"
)

func TestMcpToolCallMcpToolCallSuccessRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service_id": "value",
  "tool_call_id": "value",
  "tool_name": "value",
  "tool_description": "value",
  "parameters": {},
  "timestamp": "value",
  "state": "success",
  "result": [
    {}
  ]
}`
	assertRoundTrip(t, wire, new(McpToolCallMcpToolCallSuccess))
}
