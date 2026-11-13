package types

import (
	"testing"
)

func TestClientSecretResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value",
  "expires_at": 123,
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
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
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "tracing": "value"
  }
}`
	assertRoundTrip(t, wire, new(ClientSecretResponse))
}
