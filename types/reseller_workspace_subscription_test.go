package types

import (
	"testing"
)

func TestResellerWorkspaceSubscriptionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "billing_period": "monthly_period",
  "billing_period_end_unix": 123,
  "cancel_at_unix": 123,
  "pending_downgrade_tier": "value",
  "pending_downgrade_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ResellerWorkspaceSubscription))
}

func TestResellerWorkspaceSubscriptionRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "billing_period": "monthly_period",
  "billing_period_end_unix": 123,
  "cancel_at_unix": 123,
  "pending_downgrade_tier": "value",
  "pending_downgrade_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ResellerWorkspaceSubscription))
}
