package types

import (
	"testing"
)

func TestGuardrailTriggeredGuardrailTriggeredEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "guardrail_name": "value"
}`
	assertRoundTrip(t, wire, new(GuardrailTriggeredGuardrailTriggeredEvent))
}
