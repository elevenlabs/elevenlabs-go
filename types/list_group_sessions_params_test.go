package types

import (
	"testing"
)

func TestListGroupSessionsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_group_sessions"
}`
	assertRoundTrip(t, wire, new(ListGroupSessionsParams))
}
