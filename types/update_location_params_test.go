package types

import (
	"testing"
)

func TestUpdateLocationParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_location"
}`
	assertRoundTrip(t, wire, new(UpdateLocationParams))
}
