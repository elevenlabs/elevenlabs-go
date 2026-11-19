package types

import (
	"testing"
)

func TestModerationConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sexual": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "violence": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "violence_graphic": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "harassment": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "harassment_threatening": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "hate": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "hate_threatening": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "self_harm_instructions": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "self_harm": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "self_harm_intent": {
    "is_enabled": true,
    "threshold": 0.5
  },
  "sexual_minors": {
    "is_enabled": true,
    "threshold": 0.5
  }
}`
	assertRoundTrip(t, wire, new(ModerationConfig))
}
