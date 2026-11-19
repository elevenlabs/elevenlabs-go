package types

import (
	"testing"
)

func TestGuardrailsV1InputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": "1",
  "focus": {
    "is_enabled": true
  },
  "prompt_injection": {
    "is_enabled": true
  },
  "content": {
    "execution_mode": "streaming",
    "config": {
      "sexual": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "violence": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "harassment": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "self_harm": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "profanity": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "religion_or_politics": {
        "is_enabled": true,
        "threshold": 9.99
      },
      "medical_and_legal_information": {
        "is_enabled": true,
        "threshold": 9.99
      }
    },
    "trigger_action": {
      "type": "end_call"
    }
  },
  "custom": {
    "config": {
      "configs": [
        {
          "is_enabled": true,
          "name": "value",
          "prompt": "value",
          "execution_mode": "streaming",
          "model": "gemini-2.5-flash-lite",
          "service_tier": "standard",
          "history_message_count": 100,
          "trigger_action": {
            "type": "end_call"
          },
          "evaluate_full_response_only": true
        }
      ]
    }
  }
}`
	assertRoundTrip(t, wire, new(GuardrailsV1Input))
}
