package types

import (
	"testing"
)

func TestUpdateStaffParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_staff",
  "per_resource_timezones": true
}`
	assertRoundTrip(t, wire, new(UpdateStaffParams))
}
