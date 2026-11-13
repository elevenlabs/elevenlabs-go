package types

import (
	"testing"
)

func TestUpdateAgentProcedureParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_agent_procedure"
}`
	assertRoundTrip(t, wire, new(UpdateAgentProcedureParams))
}
