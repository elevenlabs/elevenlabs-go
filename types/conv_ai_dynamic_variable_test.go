package types

import (
	"testing"
)

func TestConvAiDynamicVariableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "variable_name": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiDynamicVariable))
}

func TestConvAiDynamicVariableRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "variable_name": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiDynamicVariable))
}

func TestConvAiDynamicVariableRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "variable_name": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiDynamicVariable))
}
