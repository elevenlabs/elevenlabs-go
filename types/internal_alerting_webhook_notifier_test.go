package types

import (
	"testing"
)

func TestInternalAlertingWebhookNotifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "webhook",
  "url": "value",
  "method": "POST",
  "headers": [
    {
      "name": "value",
      "is_secret": true,
      "value": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(InternalAlertingWebhookNotifier))
}
