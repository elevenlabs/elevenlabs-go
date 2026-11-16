package reception

import (
	"testing"
)

func TestCreatePhoneAssignmentRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "customer_facing_agent_id": "value",
  "area_codes": [
    "value"
  ],
  "number_type": "local",
  "idempotency_key": "value"
}`
	assertRoundTrip(t, wire, new(CreatePhoneAssignmentRequest))
}

func TestReassignPhoneRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "customer_facing_agent_id": "value"
}`
	assertRoundTrip(t, wire, new(ReassignPhoneRequest))
}
