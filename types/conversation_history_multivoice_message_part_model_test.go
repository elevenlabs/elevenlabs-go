package types

import (
	"testing"
)

func TestConversationHistoryMultivoiceMessagePartModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "voice_label": "value",
  "time_in_call_secs": 123
}`
	assertRoundTrip(t, wire, new(ConversationHistoryMultivoiceMessagePartModel))
}
