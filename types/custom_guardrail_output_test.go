package types

import (
	"testing"
)

func TestCustomGuardrailOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "config": {
    "configs": [
      {
        "is_enabled": true,
        "name": "value",
        "prompt": "value",
        "execution_mode": "streaming",
        "model": "gemini-2.5-flash-lite",
        "history_message_count": 100,
        "trigger_action": {
          "type": "end_call"
        },
        "evaluate_full_response_only": true
      }
    ]
  }
}`
	assertRoundTrip(t, wire, new(CustomGuardrailOutput))
}
