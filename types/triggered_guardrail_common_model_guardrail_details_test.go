package types

import (
	"testing"
)

func TestTriggeredGuardrailCommonModelGuardrailDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "rationale": "value"
}`
	assertRoundTrip(t, wire, new(TriggeredGuardrailCommonModelGuardrailDetails))
}
