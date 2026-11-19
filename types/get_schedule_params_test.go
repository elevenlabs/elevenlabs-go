package types

import (
	"testing"
)

func TestGetScheduleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_schedule",
  "include_location_filter": true
}`
	assertRoundTrip(t, wire, new(GetScheduleParams))
}
