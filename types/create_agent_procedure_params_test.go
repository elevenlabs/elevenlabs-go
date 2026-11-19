package types

import (
	"testing"
)

func TestCreateAgentProcedureParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_agent_procedure"
}`
	assertRoundTrip(t, wire, new(CreateAgentProcedureParams))
}
