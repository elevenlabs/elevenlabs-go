package types

import (
	"testing"
)

func TestPromptInjectionGuardrailRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_enabled": true
}`
	assertRoundTrip(t, wire, new(PromptInjectionGuardrail))
}
