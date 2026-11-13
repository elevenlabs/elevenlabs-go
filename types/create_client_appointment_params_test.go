package types

import (
	"testing"
)

func TestCreateClientAppointmentParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_client_appointment"
}`
	assertRoundTrip(t, wire, new(CreateClientAppointmentParams))
}
