package types

import (
	"testing"
)

func TestOwnerGetClientAppointmentsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "include_cancelled": true,
  "smb_tool_type": "owner_get_client_appointments"
}`
	assertRoundTrip(t, wire, new(OwnerGetClientAppointmentsParams))
}
