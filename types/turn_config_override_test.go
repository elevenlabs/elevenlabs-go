package types

import (
	"testing"
)

func TestTurnConfigOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "soft_timeout_config": {
    "message": "value",
    "additional_soft_timeout_messages": [
      "value"
    ]
  },
  "noisy_environment_message": "value"
}`
	assertRoundTrip(t, wire, new(TurnConfigOverride))
}

func TestTurnConfigOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "soft_timeout_config": {
    "message": "value",
    "additional_soft_timeout_messages": [
      "value"
    ]
  },
  "noisy_environment_message": "value"
}`
	assertRoundTrip(t, wire, new(TurnConfigOverride))
}
