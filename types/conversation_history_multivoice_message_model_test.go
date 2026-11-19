package types

import (
	"testing"
)

func TestConversationHistoryMultivoiceMessageModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parts": [
    {
      "text": "value",
      "voice_label": "value",
      "time_in_call_secs": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryMultivoiceMessageModel))
}

func TestConversationHistoryMultivoiceMessageModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parts": [
    {
      "text": "value",
      "voice_label": "value",
      "time_in_call_secs": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryMultivoiceMessageModel))
}

func TestConversationHistoryMultivoiceMessageModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parts": [
    {
      "text": "value",
      "voice_label": "value",
      "time_in_call_secs": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryMultivoiceMessageModel))
}
