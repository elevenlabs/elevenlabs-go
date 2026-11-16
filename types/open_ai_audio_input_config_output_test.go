package types

import (
	"testing"
)

func TestOpenAiAudioInputConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "format": {
    "type": "audio/pcm",
    "rate": 123
  },
  "transcription": {
    "language": "value"
  },
  "turn_detection": {
    "type": "semantic_vad",
    "eagerness": "low",
    "create_response": true,
    "interrupt_response": true,
    "threshold": 9.99,
    "prefix_padding_ms": 123,
    "silence_duration_ms": 123,
    "turn_timeout": 9.99
  },
  "pacing": "realtime"
}`
	assertRoundTrip(t, wire, new(OpenAiAudioInputConfigOutput))
}
