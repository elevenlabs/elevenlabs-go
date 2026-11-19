package types

import (
	"testing"
)

func TestGetAgentConfigParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_agent_config"
}`
	assertRoundTrip(t, wire, new(GetAgentConfigParams))
}
