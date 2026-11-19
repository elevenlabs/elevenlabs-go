package types

import (
	"testing"
)

func TestConversationHistoryExotelPhoneCallModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "direction": "inbound",
  "phone_number_id": "value",
  "agent_number": "value",
  "external_number": "value",
  "call_result": "answered",
  "type": "exotel",
  "stream_sid": "value",
  "call_sid": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryExotelPhoneCallModel))
}
