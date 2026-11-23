package types

import (
	"testing"
)

func TestMcpToolCallRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mcp_tool_call": {
    "service_id": "value",
    "tool_call_id": "value",
    "tool_name": "value",
    "tool_description": "value",
    "parameters": {},
    "timestamp": "value",
    "state": "loading"
  },
  "type": "mcp_tool_call"
}`
	assertRoundTrip(t, wire, new(McpToolCall))
}
