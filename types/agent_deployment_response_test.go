package types

import (
	"testing"
)

func TestAgentDeploymentResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "traffic_percentage_branch_id_map": {
    "key": 9.99
  }
}`
	assertRoundTrip(t, wire, new(AgentDeploymentResponse))
}
