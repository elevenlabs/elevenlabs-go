package types

import (
	"testing"
)

func TestWorkflowStartProcedureNodeModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "start_procedure",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "procedure_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowStartProcedureNodeModelOutput))
}
