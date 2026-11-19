package types

import (
	"testing"
)

func TestLeaveMessageParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "leave_message"
}`
	assertRoundTrip(t, wire, new(LeaveMessageParams))
}
