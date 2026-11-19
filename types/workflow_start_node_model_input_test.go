package types

import (
	"testing"
)

func TestWorkflowStartNodeModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "start",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(WorkflowStartNodeModelInput))
}
