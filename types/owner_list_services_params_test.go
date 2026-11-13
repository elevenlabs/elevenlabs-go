package types

import (
	"testing"
)

func TestOwnerListServicesParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "owner_list_services"
}`
	assertRoundTrip(t, wire, new(OwnerListServicesParams))
}
