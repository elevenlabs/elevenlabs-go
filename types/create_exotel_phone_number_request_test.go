package types

import (
	"testing"
)

func TestCreateExotelPhoneNumberRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "exotel",
  "agent_id": "value",
  "account_sid": "value",
  "api_key": "value",
  "api_token": "value",
  "api_subdomain": "api.in.exotel.com",
  "app_id": "value",
  "applet_url": "value"
}`
	assertRoundTrip(t, wire, new(CreateExotelPhoneNumberRequest))
}
