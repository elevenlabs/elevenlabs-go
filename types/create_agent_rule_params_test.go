package types

import (
	"testing"
)

func TestCreateAgentRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_agent_rule"
}`
	assertRoundTrip(t, wire, new(CreateAgentRuleParams))
}
