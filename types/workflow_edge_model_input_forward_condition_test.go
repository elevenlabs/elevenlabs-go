package types

import (
	"testing"
)

func TestWorkflowEdgeModelInputForwardConditionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "unconditional"
}`
	assertRoundTrip(t, wire, new(WorkflowEdgeModelInputForwardCondition))
}
