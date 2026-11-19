package types

import (
	"testing"
)

func TestDynamicVariableSchemaOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "dynamic_variable",
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableSchemaOverride))
}

func TestDynamicVariableSchemaOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "dynamic_variable",
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableSchemaOverride))
}

func TestDynamicVariableSchemaOverrideRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "dynamic_variable",
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableSchemaOverride))
}

func TestDynamicVariableSchemaOverrideRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "dynamic_variable",
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(DynamicVariableSchemaOverride))
}
