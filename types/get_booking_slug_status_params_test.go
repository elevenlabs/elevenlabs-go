package types

import (
	"testing"
)

func TestGetBookingSlugStatusParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_booking_slug_status"
}`
	assertRoundTrip(t, wire, new(GetBookingSlugStatusParams))
}
