package types

import (
	"testing"
)

func TestMcpToolApprovalHashRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "tool_hash": "value",
  "approval_policy": "auto_approved"
}`
	assertRoundTrip(t, wire, new(McpToolApprovalHash))
}
