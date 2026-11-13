package types

import (
	"testing"
)

func TestUnitTestGuardrailTriggerEvaluationModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "any"
}`
	assertRoundTrip(t, wire, new(UnitTestGuardrailTriggerEvaluationModel))
}

func TestUnitTestGuardrailTriggerEvaluationModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "any"
}`
	assertRoundTrip(t, wire, new(UnitTestGuardrailTriggerEvaluationModel))
}
