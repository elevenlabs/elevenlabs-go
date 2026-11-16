package types

import (
	"testing"
)

func TestWorkflowFeaturesUsageCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "tool_node": {
    "enabled": true,
    "used": true
  },
  "standalone_agent_node": {
    "enabled": true,
    "used": true
  },
  "phone_number_node": {
    "enabled": true,
    "used": true
  },
  "end_node": {
    "enabled": true,
    "used": true
  },
  "override_agent_node": {
    "enabled": true,
    "used": true
  }
}`
	assertRoundTrip(t, wire, new(WorkflowFeaturesUsageCommonModel))
}
