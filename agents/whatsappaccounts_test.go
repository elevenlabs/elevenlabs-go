package agents

import (
	"testing"
)

func TestUpdateWhatsAppAccountRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true
}`
	assertRoundTrip(t, wire, new(UpdateWhatsAppAccountRequest))
}
