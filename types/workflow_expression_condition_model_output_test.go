package types

import (
	"testing"
)

func TestWorkflowExpressionConditionModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "expression",
  "expression": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(WorkflowExpressionConditionModelOutput))
}
