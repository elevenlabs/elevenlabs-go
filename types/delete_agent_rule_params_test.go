package types

import (
	"testing"
)

func TestDeleteAgentRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_agent_rule"
}`
	assertRoundTrip(t, wire, new(DeleteAgentRuleParams))
}
