package types

import (
	"testing"
)

func TestAlertingIntegrationNotifierResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "integration",
  "integration_type": "pagerduty",
  "connection_id": "value",
  "channel_id": "value"
}`
	assertRoundTrip(t, wire, new(AlertingIntegrationNotifierResponse))
}
