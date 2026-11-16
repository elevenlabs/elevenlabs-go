package types

import (
	"testing"
)

func TestGetAppointmentByConfirmationNumberParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_appointment_by_confirmation_number"
}`
	assertRoundTrip(t, wire, new(GetAppointmentByConfirmationNumberParams))
}
