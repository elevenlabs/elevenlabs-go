package types

import (
	"testing"
)

func TestFocusGuardrailRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_enabled": true
}`
	assertRoundTrip(t, wire, new(FocusGuardrail))
}
