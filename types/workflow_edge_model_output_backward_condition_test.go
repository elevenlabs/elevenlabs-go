package types

import (
	"testing"
)

func TestWorkflowEdgeModelOutputBackwardConditionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "unconditional"
}`
	assertRoundTrip(t, wire, new(WorkflowEdgeModelOutputBackwardCondition))
}
