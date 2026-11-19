package types

import (
	"testing"
)

func TestAllowedValuesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(AllowedValues))
}

func TestAllowedValuesRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(AllowedValues))
}

func TestAllowedValuesRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dynamic_variable": "value"
}`
	assertRoundTrip(t, wire, new(AllowedValues))
}
