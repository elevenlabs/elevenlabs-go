package types

import (
	"testing"
)

func TestGetPhoneNumberExotelResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "exotel"
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberExotelResponseModel))
}

func TestGetPhoneNumberExotelResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "exotel"
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberExotelResponseModel))
}

func TestGetPhoneNumberExotelResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "exotel"
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberExotelResponseModel))
}
