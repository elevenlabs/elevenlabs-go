package types

import (
	"testing"
)

func TestUpdateGroupSessionSeatsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_group_session_seats"
}`
	assertRoundTrip(t, wire, new(UpdateGroupSessionSeatsParams))
}
