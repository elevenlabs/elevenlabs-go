package types

import (
	"testing"
)

func TestDeleteGroupSessionParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_group_session"
}`
	assertRoundTrip(t, wire, new(DeleteGroupSessionParams))
}
