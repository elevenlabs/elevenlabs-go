package reception

import (
	"testing"
)

func TestAssignResellerTierRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`
	assertRoundTrip(t, wire, new(AssignResellerTierRequest))
}

func TestChangeResellerTierRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`
	assertRoundTrip(t, wire, new(ChangeResellerTierRequest))
}
