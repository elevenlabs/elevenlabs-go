package types

import (
	"testing"
)

func TestListClientsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_clients"
}`
	assertRoundTrip(t, wire, new(ListClientsParams))
}
