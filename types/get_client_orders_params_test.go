package types

import (
	"testing"
)

func TestGetClientOrdersParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_client_orders",
  "include_cancelled": true
}`
	assertRoundTrip(t, wire, new(GetClientOrdersParams))
}
