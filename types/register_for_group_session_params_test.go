package types

import (
	"testing"
)

func TestRegisterForGroupSessionParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "register_for_group_session"
}`
	assertRoundTrip(t, wire, new(RegisterForGroupSessionParams))
}
