package types

import (
	"testing"
)

func TestDeleteCalendarEventParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_calendar_event"
}`
	assertRoundTrip(t, wire, new(DeleteCalendarEventParams))
}
