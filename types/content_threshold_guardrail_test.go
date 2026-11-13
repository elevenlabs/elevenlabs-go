package types

import (
	"testing"
)

func TestContentThresholdGuardrailRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_enabled": true,
  "threshold": 9.99
}`
	assertRoundTrip(t, wire, new(ContentThresholdGuardrail))
}
