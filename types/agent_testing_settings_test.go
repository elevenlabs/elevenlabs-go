package types

import (
	"testing"
)

func TestAgentTestingSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "attached_tests": [
    {
      "test_id": "value",
      "workflow_node_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentTestingSettings))
}
