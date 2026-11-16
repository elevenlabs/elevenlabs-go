package types

import (
	"testing"
)

func TestProvisionResellerWorkspaceResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  "has_payment_issue": true,
  "onboarding_completed_at_unix": 123,
  "member_results": [
    {
      "email": "value",
      "outcome": "invited",
      "error_status": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ProvisionResellerWorkspaceResponse))
}
