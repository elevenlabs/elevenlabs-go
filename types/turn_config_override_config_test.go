package types

import (
	"testing"
)

func TestTurnConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "soft_timeout_config": {
    "message": true,
    "additional_soft_timeout_messages": true
  },
  "noisy_environment_message": true
}`
	assertRoundTrip(t, wire, new(TurnConfigOverrideConfig))
}
