package types

import (
	"testing"
)

func TestWorkflowEndNodeModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "end",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "outcome": "success"
}`
	assertRoundTrip(t, wire, new(WorkflowEndNodeModelOutput))
}
