package workspace

import (
	"testing"
)

func TestUpdateMemberRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value",
  "is_locked": true,
  "workspace_role": "workspace_admin",
  "workspace_seat_type": "workspace_admin",
  "allow_downgrade_to_basic_seat": true
}`
	assertRoundTrip(t, wire, new(UpdateMemberRequest))
}

func TestUpdateMemberUsageLimitRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "usage_limit": 123
}`
	assertRoundTrip(t, wire, new(UpdateMemberUsageLimitRequestModel))
}
