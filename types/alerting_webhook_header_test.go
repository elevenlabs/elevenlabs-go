package types

import (
	"testing"
)

func TestAlertingWebhookHeaderRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "is_secret": true,
  "value": "value"
}`
	assertRoundTrip(t, wire, new(AlertingWebhookHeader))
}
