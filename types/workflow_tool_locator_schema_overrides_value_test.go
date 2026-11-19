package types

import (
	"testing"
)

func TestWorkflowToolLocatorSchemaOverridesValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(WorkflowToolLocatorSchemaOverridesValue))
}
