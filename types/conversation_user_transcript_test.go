package types

import (
	"testing"
)

func TestConversationUserTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_transcription_event": {
    "user_transcript": "value",
    "event_id": 123
  },
  "type": "user_transcript"
}`
	assertRoundTrip(t, wire, new(ConversationUserTranscript))
}
