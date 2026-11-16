package types

import (
	"testing"
)

func TestGuardrailLlmVerdictRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "rationale": "value"
}`
	assertRoundTrip(t, wire, new(GuardrailLlmVerdict))
}
