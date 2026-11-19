package types

import (
	"testing"
)

func TestWorkflowToolNodeModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "tool",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "tools": [
    {
      "tool_id": "value",
      "schema_overrides": {
        "key": {
          "source": "constant",
          "constant_value": "value"
        }
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(WorkflowToolNodeModelOutput))
}
