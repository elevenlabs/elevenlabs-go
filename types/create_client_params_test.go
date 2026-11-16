package types

import (
	"testing"
)

func TestCreateClientParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_client"
}`
	assertRoundTrip(t, wire, new(CreateClientParams))
}
