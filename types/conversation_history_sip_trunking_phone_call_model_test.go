package types

import (
	"testing"
)

func TestConversationHistorySipTrunkingPhoneCallModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "direction": "inbound",
  "phone_number_id": "value",
  "agent_number": "value",
  "external_number": "value",
  "call_result": "answered",
  "type": "sip_trunking",
  "call_id": "value",
  "call_sid": "value",
  "sip_header_dynamic_variables": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistorySipTrunkingPhoneCallModel))
}
