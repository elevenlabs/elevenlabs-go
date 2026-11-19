package types

import (
	"testing"
)

func TestGuardrailsResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "triggered": true
}`
	assertRoundTrip(t, wire, new(GuardrailsResult))
}
