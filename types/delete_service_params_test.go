package types

import (
	"testing"
)

func TestDeleteServiceParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_service"
}`
	assertRoundTrip(t, wire, new(DeleteServiceParams))
}
