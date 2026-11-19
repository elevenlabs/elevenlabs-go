package types

import (
	"testing"
)

func TestGuardrailUnknownVerdictRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "unknown"
}`
	assertRoundTrip(t, wire, new(GuardrailUnknownVerdict))
}
