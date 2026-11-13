package types

import (
	"testing"
)

func TestConversationTtsUsageModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "primary_tts_model": "value",
  "total_audio_output_seconds": 9.99,
  "total_characters": 123,
  "per_voice_usage": [
    {
      "voice_id": "value",
      "audio_output_seconds": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(ConversationTtsUsageModel))
}
