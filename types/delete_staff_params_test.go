package types

import (
	"testing"
)

func TestDeleteStaffParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_staff"
}`
	assertRoundTrip(t, wire, new(DeleteStaffParams))
}
