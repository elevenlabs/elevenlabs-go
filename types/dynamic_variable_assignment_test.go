package types

import (
	"testing"
)

func TestDynamicVariableAssignmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "response",
  "dynamic_variable": "value",
  "value_path": "value",
  "sanitize": true,
  "preserve_native_type": true
}`
	assertRoundTrip(t, wire, new(DynamicVariableAssignment))
}

func TestDynamicVariableAssignmentRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "response",
  "dynamic_variable": "value",
  "value_path": "value",
  "sanitize": true,
  "preserve_native_type": true
}`
	assertRoundTrip(t, wire, new(DynamicVariableAssignment))
}

func TestDynamicVariableAssignmentRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "response",
  "dynamic_variable": "value",
  "value_path": "value",
  "sanitize": true,
  "preserve_native_type": true
}`
	assertRoundTrip(t, wire, new(DynamicVariableAssignment))
}
