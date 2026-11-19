package types

import (
	"testing"
)

func TestSmsConversationInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "direction": "inbound",
  "phone_number_id": "value",
  "sms_user_phone_number": "value",
  "agent_phone_number": "value"
}`
	assertRoundTrip(t, wire, new(SmsConversationInfo))
}
