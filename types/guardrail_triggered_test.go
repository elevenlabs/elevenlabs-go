package types

import (
	"testing"
)

func TestGuardrailTriggeredRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "guardrail_triggered_event": {
    "guardrail_name": "value"
  },
  "type": "guardrail_triggered"
}`
	assertRoundTrip(t, wire, new(GuardrailTriggered))
}
