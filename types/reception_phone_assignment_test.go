package types

import (
	"testing"
)

func TestReceptionPhoneAssignmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`
	assertRoundTrip(t, wire, new(ReceptionPhoneAssignment))
}
