package mcpservers

import (
	"testing"
)

func TestMcpToolAddApprovalRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "tool_description": "value",
  "input_schema": {
    "key": {}
  },
  "approval_policy": "auto_approved"
}`
	assertRoundTrip(t, wire, new(McpToolAddApprovalRequestModel))
}
