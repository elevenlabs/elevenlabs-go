package types

import (
	"testing"
)

func TestListServicesParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "list_services"
}`
	assertRoundTrip(t, wire, new(ListServicesParams))
}
