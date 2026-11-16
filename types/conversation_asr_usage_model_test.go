package types

import (
	"testing"
)

func TestConversationAsrUsageModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "asr_model": "value",
  "total_transcription_calls": 123,
  "total_audio_input_seconds": 9.99
}`
	assertRoundTrip(t, wire, new(ConversationAsrUsageModel))
}
