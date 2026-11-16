package types

import (
	"testing"
)

func TestWorkflowToolLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_id": "value",
  "schema_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(WorkflowToolLocator))
}
