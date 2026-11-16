package types

import (
	"testing"
)

func TestConstantSchemaOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(ConstantSchemaOverride))
}

func TestConstantSchemaOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(ConstantSchemaOverride))
}

func TestConstantSchemaOverrideRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(ConstantSchemaOverride))
}
