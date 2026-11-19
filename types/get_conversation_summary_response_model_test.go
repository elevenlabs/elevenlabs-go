package types

import (
	"testing"
)

func TestGetConversationSummaryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "agent_id": "value",
  "status": "initiated",
  "call_summary_title": "value",
  "transcript_summary": "value",
  "call_successful": "success",
  "message_count": 123,
  "messages": [
    {
      "role": "user",
      "message": "value"
    }
  ],
  "messages_omitted": true,
  "note": "value"
}`
	assertRoundTrip(t, wire, new(GetConversationSummaryResponseModel))
}
