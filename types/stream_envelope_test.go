package types

import (
	"testing"
)

func TestStreamEnvelopeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "turn_id": "value",
  "created_at_unix_ms": 123,
  "event": {
    "type": "stream_ended",
    "stream_ended_event": {
      "status": "in_progress",
      "error": {
        "type": "value",
        "code": "value",
        "message": "value"
      },
      "required_actions": [
        {
          "type": "client_tool_call",
          "tool_call_id": "value",
          "tool_name": "value",
          "parameters": {
            "key": {}
          }
        }
      ]
    }
  }
}`
	assertRoundTrip(t, wire, new(StreamEnvelope))
}
