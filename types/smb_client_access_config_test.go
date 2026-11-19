package types

import (
	"testing"
)

func TestSmbClientAccessConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enable_secure_mode": true,
  "identity_verification_enabled": true,
  "identity_verification_email_enabled": true,
  "identity_verification_sms_enabled": true
}`
	assertRoundTrip(t, wire, new(SmbClientAccessConfig))
}
