package types

import (
	"testing"
)

func TestUpdateClientParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_client"
}`
	assertRoundTrip(t, wire, new(UpdateClientParams))
}
