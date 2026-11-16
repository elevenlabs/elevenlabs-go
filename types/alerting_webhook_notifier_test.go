package types

import (
	"testing"
)

func TestAlertingWebhookNotifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "webhook",
  "webhook_id": "value"
}`
	assertRoundTrip(t, wire, new(AlertingWebhookNotifier))
}
