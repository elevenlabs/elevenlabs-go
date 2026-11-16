package types

import (
	"testing"
)

func TestPendingSubscriptionSwitchResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "change",
  "next_tier": "free",
  "next_billing_period": "monthly_period",
  "timestamp_seconds": 123
}`
	assertRoundTrip(t, wire, new(PendingSubscriptionSwitchResponseModel))
}
