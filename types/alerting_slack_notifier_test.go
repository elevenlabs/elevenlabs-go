package types

import (
	"testing"
)

func TestAlertingSlackNotifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "integration",
  "integration_type": "slack",
  "connection_id": "value",
  "channel_id": "value"
}`
	assertRoundTrip(t, wire, new(AlertingSlackNotifier))
}
