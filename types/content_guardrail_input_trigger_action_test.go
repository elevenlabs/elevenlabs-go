package types

import (
	"testing"
)

func TestContentGuardrailInputTriggerActionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "end_call"
}`
	assertRoundTrip(t, wire, new(ContentGuardrailInputTriggerAction))
}
