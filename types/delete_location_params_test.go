package types

import (
	"testing"
)

func TestDeleteLocationParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_location"
}`
	assertRoundTrip(t, wire, new(DeleteLocationParams))
}
