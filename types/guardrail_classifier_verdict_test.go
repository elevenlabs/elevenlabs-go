package types

import (
	"testing"
)

func TestGuardrailClassifierVerdictRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "classifier",
  "score": 9.99,
  "threshold": 9.99
}`
	assertRoundTrip(t, wire, new(GuardrailClassifierVerdict))
}
