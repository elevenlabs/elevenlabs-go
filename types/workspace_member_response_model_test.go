package types

import (
	"testing"
)

func TestWorkspaceMemberResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_id": "value",
  "email": "value",
  "first_name": "value",
  "seat_type": "workspace_admin",
  "is_owner": true,
  "is_locked": true
}`
	assertRoundTrip(t, wire, new(WorkspaceMemberResponseModel))
}
