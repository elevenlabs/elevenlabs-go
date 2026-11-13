package types

import (
	"testing"
)

func TestUpdateHolidayParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_holiday"
}`
	assertRoundTrip(t, wire, new(UpdateHolidayParams))
}
