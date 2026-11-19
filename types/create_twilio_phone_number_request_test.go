package types

import (
	"testing"
)

func TestCreateTwilioPhoneNumberRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "twilio",
  "agent_id": "value",
  "sid": "value",
  "token": "value",
  "account_auth_token": "value",
  "region_config": {
    "region_id": "us1",
    "token": "value",
    "edge_location": "ashburn"
  },
  "enable_sms": true
}`
	assertRoundTrip(t, wire, new(CreateTwilioPhoneNumberRequest))
}
