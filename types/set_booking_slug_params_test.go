package types

import (
	"testing"
)

func TestSetBookingSlugParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "set_booking_slug"
}`
	assertRoundTrip(t, wire, new(SetBookingSlugParams))
}
