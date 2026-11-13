package types

import (
	"testing"
)

func TestListAgentProceduresParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_agent_procedures"
}`
	assertRoundTrip(t, wire, new(ListAgentProceduresParams))
}
