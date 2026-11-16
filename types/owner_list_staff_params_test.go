package types

import (
	"testing"
)

func TestOwnerListStaffParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "owner_list_staff"
}`
	assertRoundTrip(t, wire, new(OwnerListStaffParams))
}
