package types

import (
	"testing"
)

func TestWorkflowStateUpdateInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "dynamic_variable",
  "variable_name": "value",
  "expression": {
    "type": "string_literal",
    "value": "value"
  },
  "required": true
}`
	assertRoundTrip(t, wire, new(WorkflowStateUpdateInput))
}
