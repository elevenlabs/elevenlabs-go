package types

import (
	"testing"
)

func TestAlertingPagerDutyNotifierResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "integration",
  "integration_type": "pagerduty",
  "connection_id": "value"
}`
	assertRoundTrip(t, wire, new(AlertingPagerDutyNotifierResponse))
}
