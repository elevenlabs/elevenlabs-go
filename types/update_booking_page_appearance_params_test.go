package types

import (
	"testing"
)

func TestUpdateBookingPageAppearanceParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_booking_page_appearance"
}`
	assertRoundTrip(t, wire, new(UpdateBookingPageAppearanceParams))
}
