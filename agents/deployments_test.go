package agents

import (
	"testing"
)

func TestBodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "deployment_request": {
    "requests": [
      {
        "branch_id": "value",
        "deployment_strategy": {
          "type": "percentage",
          "traffic_percentage": 9.99
        }
      }
    ]
  }
}`
	assertRoundTrip(t, wire, new(BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost))
}
