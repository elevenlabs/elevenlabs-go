package types

import (
	"testing"
)

func TestConversationUserTranscriptUserTranscriptionEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_transcript": "value",
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(ConversationUserTranscriptUserTranscriptionEvent))
}
