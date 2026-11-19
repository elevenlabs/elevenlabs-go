package types

import (
	"testing"
)

func TestConversationHistoryTwilioPhoneCallModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "direction": "inbound",
  "phone_number_id": "value",
  "agent_number": "value",
  "external_number": "value",
  "call_result": "answered",
  "type": "twilio",
  "stream_sid": "value",
  "call_sid": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTwilioPhoneCallModel))
}
