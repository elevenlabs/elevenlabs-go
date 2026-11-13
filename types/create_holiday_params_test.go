package types

import (
	"testing"
)

func TestCreateHolidayParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_holiday"
}`
	assertRoundTrip(t, wire, new(CreateHolidayParams))
}
