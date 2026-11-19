package types

import (
	"testing"
)

func TestMcpApprovalPolicyUpdateRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "approval_policy": "auto_approve_all"
}`
	assertRoundTrip(t, wire, new(McpApprovalPolicyUpdateRequestModel))
}
