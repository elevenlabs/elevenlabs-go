package types

import (
	"testing"
)

func TestTwilioMachineDetectionConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "enable",
  "actions": [
    {
      "action": "hang_up",
      "answered_by": [
        "human"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(TwilioMachineDetectionConfig))
}

func TestTwilioMachineDetectionConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "enable",
  "actions": [
    {
      "action": "hang_up",
      "answered_by": [
        "human"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(TwilioMachineDetectionConfig))
}

func TestTwilioMachineDetectionConfigRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "enable",
  "actions": [
    {
      "action": "hang_up",
      "answered_by": [
        "human"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(TwilioMachineDetectionConfig))
}
