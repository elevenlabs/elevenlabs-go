package types

import (
	"testing"
)

func TestResellerUpcomingPaymentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "workspace_id": "value",
  "workspace_name": "value",
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period",
    "billing_period_end_unix": 123,
    "cancel_at_unix": 123,
    "pending_downgrade_tier": "value",
    "pending_downgrade_at_unix": 123
  },
  "plan_usd": 9.99,
  "overage_so_far_usd": 9.99
}`
	assertRoundTrip(t, wire, new(ResellerUpcomingPayment))
}
