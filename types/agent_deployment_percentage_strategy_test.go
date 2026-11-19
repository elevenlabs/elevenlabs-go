package types

import (
	"testing"
)

func TestAgentDeploymentPercentageStrategyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "percentage",
  "traffic_percentage": 9.99
}`
	assertRoundTrip(t, wire, new(AgentDeploymentPercentageStrategy))
}
