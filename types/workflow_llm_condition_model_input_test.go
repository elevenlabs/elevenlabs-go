package types

import (
	"testing"
)

func TestWorkflowLlmConditionModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "type": "llm",
  "condition": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowLlmConditionModelInput))
}
