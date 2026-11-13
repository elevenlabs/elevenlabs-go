package types

import (
	"testing"
)

func TestConversationHistoryTranscriptToolCallClientDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client",
  "parameters": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallClientDetails))
}

func TestConversationHistoryTranscriptToolCallClientDetailsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client",
  "parameters": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallClientDetails))
}

func TestConversationHistoryTranscriptToolCallClientDetailsRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client",
  "parameters": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallClientDetails))
}
