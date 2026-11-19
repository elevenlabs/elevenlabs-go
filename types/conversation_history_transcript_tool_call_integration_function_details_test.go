package types

import (
	"testing"
)

func TestConversationHistoryTranscriptToolCallIntegrationFunctionDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resolved_parameters": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallIntegrationFunctionDetails))
}

func TestConversationHistoryTranscriptToolCallIntegrationFunctionDetailsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resolved_parameters": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallIntegrationFunctionDetails))
}
