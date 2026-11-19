package types

import (
	"testing"
)

func TestCancelGroupSessionRegistrationParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "cancel_group_session_registration"
}`
	assertRoundTrip(t, wire, new(CancelGroupSessionRegistrationParams))
}
