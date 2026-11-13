package types

import (
	"testing"
)

func TestDeleteHolidayParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_holiday"
}`
	assertRoundTrip(t, wire, new(DeleteHolidayParams))
}
