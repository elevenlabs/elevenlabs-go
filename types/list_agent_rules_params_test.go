package types

import (
	"testing"
)

func TestListAgentRulesParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_agent_rules"
}`
	assertRoundTrip(t, wire, new(ListAgentRulesParams))
}
