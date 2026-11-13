package types

import (
	"testing"
)

func TestWorkflowUnconditionalModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "unconditional"
}`
	assertRoundTrip(t, wire, new(WorkflowUnconditionalModelInput))
}
