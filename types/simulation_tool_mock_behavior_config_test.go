package types

import (
	"testing"
)

func TestSimulationToolMockBehaviorConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mocking_strategy": "all",
  "fallback_strategy": "call_real_tool",
  "mocked_tool_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SimulationToolMockBehaviorConfig))
}

func TestSimulationToolMockBehaviorConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mocking_strategy": "all",
  "fallback_strategy": "call_real_tool",
  "mocked_tool_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SimulationToolMockBehaviorConfig))
}
