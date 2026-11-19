package types

import (
	"testing"
)

func TestWorkflowUpdateStateNodeModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "update_state",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "label": "value",
  "updates": [
    {
      "type": "dynamic_variable",
      "variable_name": "value",
      "expression": {
        "type": "string_literal",
        "value": "value"
      },
      "required": true
    }
  ]
}`
	assertRoundTrip(t, wire, new(WorkflowUpdateStateNodeModelInput))
}
