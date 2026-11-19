package types

import (
	"testing"
)

func TestGetBookingPageSettingsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_booking_page_settings"
}`
	assertRoundTrip(t, wire, new(GetBookingPageSettingsParams))
}
