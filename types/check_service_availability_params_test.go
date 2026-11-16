package types

import (
	"testing"
)

func TestCheckServiceAvailabilityParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "check_service_availability"
}`
	assertRoundTrip(t, wire, new(CheckServiceAvailabilityParams))
}
