package types

import (
	"testing"
)

func TestTriggeredGuardrailCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "guardrail_type": "custom",
  "guardrail_name": "value",
  "execution_mode": "streaming",
  "guardrail_details": {
    "type": "llm",
    "rationale": "value"
  }
}`
	assertRoundTrip(t, wire, new(TriggeredGuardrailCommonModel))
}

func TestTriggeredGuardrailCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "guardrail_type": "custom",
  "guardrail_name": "value",
  "execution_mode": "streaming",
  "guardrail_details": {
    "type": "llm",
    "rationale": "value"
  }
}`
	assertRoundTrip(t, wire, new(TriggeredGuardrailCommonModel))
}

func TestTriggeredGuardrailCommonModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "guardrail_type": "custom",
  "guardrail_name": "value",
  "execution_mode": "streaming",
  "guardrail_details": {
    "type": "llm",
    "rationale": "value"
  }
}`
	assertRoundTrip(t, wire, new(TriggeredGuardrailCommonModel))
}
