package types

import (
	"testing"
)

func TestOwnerListLocationsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "owner_list_locations"
}`
	assertRoundTrip(t, wire, new(OwnerListLocationsParams))
}
