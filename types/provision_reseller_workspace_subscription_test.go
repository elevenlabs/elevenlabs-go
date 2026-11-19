package types

import (
	"testing"
)

func TestProvisionResellerWorkspaceSubscriptionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "billing_period": "monthly_period"
}`
	assertRoundTrip(t, wire, new(ProvisionResellerWorkspaceSubscription))
}
