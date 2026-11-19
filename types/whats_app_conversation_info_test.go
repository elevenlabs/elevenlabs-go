package types

import (
	"testing"
)

func TestWhatsAppConversationInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "direction": "inbound",
  "whatsapp_phone_number_id": "value",
  "whatsapp_user_id": "value",
  "awaiting_first_user_message": true,
  "recipient": {
    "phone": "value",
    "bsuid": "value"
  },
  "call_permission_request": {
    "state": "blocked",
    "requested_at_unix": 123,
    "responded_at_unix": 123,
    "permanent": true,
    "granted_by_user_call": true
  },
  "call_result": "answered",
  "opening_message": {
    "template_name": "value",
    "wamid": "value",
    "sent_at_unix": 123,
    "delivered_at_unix": 123,
    "read_at_unix": 123,
    "replied_at_unix": 123,
    "failed_at_unix": 123,
    "error_code": "value",
    "pricing_category": "value",
    "billable": true
  }
}`
	assertRoundTrip(t, wire, new(WhatsAppConversationInfo))
}
