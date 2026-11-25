package types

import (
	"testing"
)

func TestResellerWorkspacesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "workspaces": [
    {
      "workspace_id": "value",
      "business_name": "value",
      "billing_tier": "value",
      "receptionist_count": 123,
      "calls_last_30_days": 123,
      "created_at_unix": 123,
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "charge_customer": true,
      "has_payment_issue": true,
      "onboarding_completed_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(ResellerWorkspacesResponse))
}
