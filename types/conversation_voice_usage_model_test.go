package types

import (
	"testing"
)

func TestConversationVoiceUsageModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "audio_output_seconds": 9.99
}`
	assertRoundTrip(t, wire, new(ConversationVoiceUsageModel))
}
