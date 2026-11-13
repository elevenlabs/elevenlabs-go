package types

import (
	"testing"
)

func TestListStaffParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "list_staff"
}`
	assertRoundTrip(t, wire, new(ListStaffParams))
}
