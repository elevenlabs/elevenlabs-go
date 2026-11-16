package types

import (
	"testing"
)

func TestReceptionTeamMemberResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_id": "value",
  "email": "value",
  "first_name": "value",
  "is_owner": true,
  "is_admin": true,
  "picture_url": "value",
  "is_suspended": true
}`
	assertRoundTrip(t, wire, new(ReceptionTeamMemberResponse))
}
