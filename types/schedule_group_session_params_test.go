package types

import (
	"testing"
)

func TestScheduleGroupSessionParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "schedule_group_session"
}`
	assertRoundTrip(t, wire, new(ScheduleGroupSessionParams))
}
