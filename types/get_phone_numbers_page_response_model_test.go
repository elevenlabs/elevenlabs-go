package types

import (
	"testing"
)

func TestGetPhoneNumbersPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_numbers": [
    {
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
      "provider": "twilio"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetPhoneNumbersPageResponseModel))
}
