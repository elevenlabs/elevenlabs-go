package types

import (
	"testing"
)

func TestCancelCalendarEventParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "cancel_calendar_event"
}`
	assertRoundTrip(t, wire, new(CancelCalendarEventParams))
}
