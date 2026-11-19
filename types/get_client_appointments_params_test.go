package types

import (
	"testing"
)

func TestGetClientAppointmentsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "include_cancelled": true,
  "smb_tool_type": "get_client_appointments"
}`
	assertRoundTrip(t, wire, new(GetClientAppointmentsParams))
}
