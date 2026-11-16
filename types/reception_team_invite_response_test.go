package types

import (
	"testing"
)

func TestReceptionTeamInviteResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value",
  "seat_type": "workspace_admin",
  "inviting_user_email": "value",
  "created_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionTeamInviteResponse))
}
