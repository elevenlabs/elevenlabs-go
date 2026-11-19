package types

import (
	"testing"
)

func TestListLocationsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_locations"
}`
	assertRoundTrip(t, wire, new(ListLocationsParams))
}
