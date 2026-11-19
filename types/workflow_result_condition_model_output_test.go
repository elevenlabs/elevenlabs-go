package types

import (
	"testing"
)

func TestWorkflowResultConditionModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "result",
  "successful": true
}`
	assertRoundTrip(t, wire, new(WorkflowResultConditionModelOutput))
}
