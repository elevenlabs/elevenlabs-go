package types

import (
	"testing"
)

func TestAlertingSlackNotifierResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "integration",
  "integration_type": "slack",
  "connection_id": "value",
  "channel_id": "value"
}`
	assertRoundTrip(t, wire, new(AlertingSlackNotifierResponse))
}
