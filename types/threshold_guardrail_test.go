package types

import (
	"testing"
)

func TestThresholdGuardrailRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_enabled": true,
  "threshold": 0.5
}`
	assertRoundTrip(t, wire, new(ThresholdGuardrail))
}
