package types

import (
	"testing"
)

func TestUpdateAgentRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_agent_rule"
}`
	assertRoundTrip(t, wire, new(UpdateAgentRuleParams))
}
