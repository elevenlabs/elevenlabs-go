package types

import (
	"testing"
)

func TestWorkflowEdgeModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "value",
  "target": "value",
  "forward_condition": {
    "label": "value",
    "type": "unconditional"
  },
  "backward_condition": {
    "label": "value",
    "type": "unconditional"
  },
  "reroutes": [
    {
      "x": 9.99,
      "y": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(WorkflowEdgeModelOutput))
}
