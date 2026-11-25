package reception

import (
	"testing"
)

func TestReceptionProvisionResellerWorkspaceRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_name": "value",
  "vertical": "salon_barbershop",
  "charge_customer": true,
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
	assertRoundTrip(t, wire, new(ReceptionProvisionResellerWorkspaceRequest))
}
