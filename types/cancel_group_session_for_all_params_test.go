package types

import (
	"testing"
)

func TestCancelGroupSessionForAllParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "cancel_group_session_for_all"
}`
	assertRoundTrip(t, wire, new(CancelGroupSessionForAllParams))
}
