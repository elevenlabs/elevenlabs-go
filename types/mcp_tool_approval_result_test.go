package types

import (
	"testing"
)

func TestMcpToolApprovalResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp_tool_approval_result",
  "tool_call_id": "value",
  "is_approved": true
}`
	assertRoundTrip(t, wire, new(McpToolApprovalResult))
}
