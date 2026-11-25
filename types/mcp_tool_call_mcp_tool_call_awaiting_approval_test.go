package types

import (
	"testing"
)

func TestMcpToolCallMcpToolCallAwaitingApprovalRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service_id": "value",
  "tool_call_id": "value",
  "tool_name": "value",
  "tool_description": "value",
  "parameters": {
    "key": {}
  },
  "timestamp": "value",
  "state": "awaiting_approval",
  "approval_timeout_secs": 123
}`
	assertRoundTrip(t, wire, new(McpToolCallMcpToolCallAwaitingApproval))
}
