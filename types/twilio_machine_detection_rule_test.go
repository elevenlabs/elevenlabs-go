package types

import (
	"testing"
)

func TestTwilioMachineDetectionRuleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "action": "hang_up",
  "answered_by": [
    "human"
  ]
}`
	assertRoundTrip(t, wire, new(TwilioMachineDetectionRule))
}
