package types

import (
	"testing"
)

func TestListCalendarEventsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_calendar_events",
  "include_cancelled": true
}`
	assertRoundTrip(t, wire, new(ListCalendarEventsParams))
}
