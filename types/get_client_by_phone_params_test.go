package types

import (
	"testing"
)

func TestGetClientByPhoneParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_client_by_phone"
}`
	assertRoundTrip(t, wire, new(GetClientByPhoneParams))
}
