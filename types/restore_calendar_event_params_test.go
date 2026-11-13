package types

import (
	"testing"
)

func TestRestoreCalendarEventParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "restore_calendar_event"
}`
	assertRoundTrip(t, wire, new(RestoreCalendarEventParams))
}
