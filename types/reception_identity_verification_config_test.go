package types

import (
	"testing"
)

func TestReceptionIdentityVerificationConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "channels": [
    "email"
  ],
  "sms_phone_number_id": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionIdentityVerificationConfig))
}
