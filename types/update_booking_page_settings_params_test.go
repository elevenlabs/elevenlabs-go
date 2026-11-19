package types

import (
	"testing"
)

func TestUpdateBookingPageSettingsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_booking_page_settings"
}`
	assertRoundTrip(t, wire, new(UpdateBookingPageSettingsParams))
}
