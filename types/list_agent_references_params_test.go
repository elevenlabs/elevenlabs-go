package types

import (
	"testing"
)

func TestListAgentReferencesParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_agent_references"
}`
	assertRoundTrip(t, wire, new(ListAgentReferencesParams))
}
