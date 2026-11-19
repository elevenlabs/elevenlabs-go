package types

import (
	"testing"
)

func TestUpdateCalendarEventParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_calendar_event"
}`
	assertRoundTrip(t, wire, new(UpdateCalendarEventParams))
}
