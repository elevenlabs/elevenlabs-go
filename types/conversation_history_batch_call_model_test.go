package types

import (
	"testing"
)

func TestConversationHistoryBatchCallModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "batch_call_id": "value",
  "batch_call_recipient_id": "value",
  "campaign": {
    "campaign_id": "value",
    "campaign_lead_id": "value",
    "cold_attempt": 123,
    "conversation_type": "initial"
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistoryBatchCallModel))
}
