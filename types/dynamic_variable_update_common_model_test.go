package types

import (
	"testing"
)

func TestDynamicVariableUpdateCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "variable_name": "value",
  "old_value": "value",
  "new_value": "value",
  "updated_at": 9.99,
  "tool_name": "value",
  "tool_request_id": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableUpdateCommonModel))
}

func TestDynamicVariableUpdateCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "variable_name": "value",
  "old_value": "value",
  "new_value": "value",
  "updated_at": 9.99,
  "tool_name": "value",
  "tool_request_id": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableUpdateCommonModel))
}
