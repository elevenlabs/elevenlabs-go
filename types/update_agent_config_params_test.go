package types

import (
	"testing"
)

func TestUpdateAgentConfigParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_agent_config"
}`
	assertRoundTrip(t, wire, new(UpdateAgentConfigParams))
}
