package types

import (
	"testing"
)

func TestAlertingWebhookSecretHeaderValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "encrypted_value": "value",
  "nonce": "value"
}`
	assertRoundTrip(t, wire, new(AlertingWebhookSecretHeaderValue))
}
