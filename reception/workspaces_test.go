package reception

import (
	"testing"
)

func TestProvisionResellerWorkspaceRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_name": "value",
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period"
  },
  "phone_number": "value",
  "members": [
    {
      "email": "value",
      "role": "workspace_admin"
    }
  ],
  "complete_onboarding": true,
  "idempotency_key": "value"
}`
	assertRoundTrip(t, wire, new(ProvisionResellerWorkspaceRequest))
}
