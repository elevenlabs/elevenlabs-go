package types

import (
	"testing"
)

func TestTeamInviteResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "email": "value",
  "outcome": "invited",
  "error_status": "value"
}`
	assertRoundTrip(t, wire, new(TeamInviteResult))
}
