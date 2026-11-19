package types

import (
	"testing"
)

func TestAgentDeploymentRequestItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_id": "value",
  "deployment_strategy": {
    "type": "percentage",
    "traffic_percentage": 9.99
  }
}`
	assertRoundTrip(t, wire, new(AgentDeploymentRequestItem))
}
