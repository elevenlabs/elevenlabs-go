package types

import (
	"testing"
)

func TestConversationTurnMetricsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metrics": {
    "key": {
      "elapsed_time": 9.99
    }
  },
  "convai_asr_provider": "value",
  "convai_tts_model": "value",
  "convai_tts_cascade": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ConversationTurnMetrics))
}

func TestConversationTurnMetricsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metrics": {
    "key": {
      "elapsed_time": 9.99
    }
  },
  "convai_asr_provider": "value",
  "convai_tts_model": "value",
  "convai_tts_cascade": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ConversationTurnMetrics))
}

func TestConversationTurnMetricsRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metrics": {
    "key": {
      "elapsed_time": 9.99
    }
  },
  "convai_asr_provider": "value",
  "convai_tts_model": "value",
  "convai_tts_cascade": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ConversationTurnMetrics))
}
