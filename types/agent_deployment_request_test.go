package types

import (
	"testing"
)

func TestAgentDeploymentRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "requests": [
    {
      "branch_id": "value",
      "deployment_strategy": {
        "type": "percentage",
        "traffic_percentage": 9.99
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentDeploymentRequest))
}
