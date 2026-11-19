package types

import (
	"testing"
)

func TestOrchestratorToolMockBehaviorConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mocking_strategy": "all",
  "fallback_strategy": "call_real_tool",
  "mocked_tool_names": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(OrchestratorToolMockBehaviorConfig))
}
