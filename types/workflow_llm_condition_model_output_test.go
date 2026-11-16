package types

import (
	"testing"
)

func TestWorkflowLlmConditionModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "llm",
  "condition": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowLlmConditionModelOutput))
}
