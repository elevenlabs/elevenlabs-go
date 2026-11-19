package types

import (
	"testing"
)

func TestDeleteAgentProcedureParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_agent_procedure"
}`
	assertRoundTrip(t, wire, new(DeleteAgentProcedureParams))
}
