package types

import (
	"testing"
)

func TestDeleteClientParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_client"
}`
	assertRoundTrip(t, wire, new(DeleteClientParams))
}
