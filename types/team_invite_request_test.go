package types

import (
	"testing"
)

func TestTeamInviteRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value",
  "role": "workspace_admin"
}`
	assertRoundTrip(t, wire, new(TeamInviteRequest))
}
