package types

import (
	"testing"
)

func TestCreateStaffParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_staff",
  "per_resource_timezones": true
}`
	assertRoundTrip(t, wire, new(CreateStaffParams))
}
