package types

import (
	"testing"
)

func TestSafetyResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_blocked_ivc": true,
  "is_blocked_non_ivc": true,
  "ignore_safety_evaluation": true
}`
	assertRoundTrip(t, wire, new(SafetyResponseModel))
}
