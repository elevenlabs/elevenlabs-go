package types

import (
	"testing"
)

func TestPhoneVoiceConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "ringing_timeout_secs": 123,
  "twilio_call_recording_enabled": true,
  "twilio_machine_detection": {
    "mode": "enable",
    "actions": [
      {
        "action": "hang_up",
        "answered_by": [
          "human"
        ]
      }
    ]
  },
  "mode": "voice"
}`
	assertRoundTrip(t, wire, new(PhoneVoiceConfigInput))
}
