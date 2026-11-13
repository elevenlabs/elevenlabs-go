package types

import (
	"testing"
)

func TestUnitTestToolCallEvaluationModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parameters": [
    {
      "eval": {
        "type": "llm",
        "description": "value"
      },
      "path": "value"
    }
  ],
  "referenced_tool": {
    "id": "value",
    "type": "system"
  },
  "verify_absence": true,
  "workflow_node_transition": {
    "type": "node_id",
    "agent_id": "value",
    "target_node_id": "value"
  },
  "guardrail_trigger": {
    "type": "any"
  }
}`
	assertRoundTrip(t, wire, new(UnitTestToolCallEvaluationModelInput))
}
