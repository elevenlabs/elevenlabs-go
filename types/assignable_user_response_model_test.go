package types

import (
	"testing"
)

func TestAssignableUserResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_id": "value",
  "email": "value",
  "first_name": "value",
  "is_service_account": true,
  "has_access": true
}`
	assertRoundTrip(t, wire, new(AssignableUserResponseModel))
}
