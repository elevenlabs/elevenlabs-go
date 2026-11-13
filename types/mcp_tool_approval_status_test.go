package types

import (
	"testing"
)

func TestMcpToolApprovalStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_id": "value",
  "state": "up_to_date",
  "approval_policy": "auto_approved",
  "approved_definition": {
    "description": "value",
    "input_schema": {
      "key": {}
    }
  }
}`
	assertRoundTrip(t, wire, new(McpToolApprovalStatus))
}
