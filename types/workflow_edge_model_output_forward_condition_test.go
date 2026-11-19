package types

import (
	"testing"
)

func TestWorkflowEdgeModelOutputForwardConditionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "unconditional"
}`
	assertRoundTrip(t, wire, new(WorkflowEdgeModelOutputForwardCondition))
}
