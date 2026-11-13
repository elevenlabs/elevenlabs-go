package types

import (
	"testing"
)

func TestWebhookHmacSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_type": "hmac",
  "name": "value",
  "webhook_url": "value",
  "request_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(WebhookHmacSettings))
}
