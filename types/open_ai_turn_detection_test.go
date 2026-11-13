package types

import (
	"testing"
)

func TestOpenAiTurnDetectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "semantic_vad",
  "eagerness": "low",
  "create_response": true,
  "interrupt_response": true,
  "threshold": 9.99,
  "prefix_padding_ms": 123,
  "silence_duration_ms": 123,
  "turn_timeout": 9.99
}`
	assertRoundTrip(t, wire, new(OpenAiTurnDetection))
}
