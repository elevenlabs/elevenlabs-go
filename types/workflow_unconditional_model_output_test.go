package types

import (
	"testing"
)

func TestWorkflowUnconditionalModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "unconditional"
}`
	assertRoundTrip(t, wire, new(WorkflowUnconditionalModelOutput))
}
