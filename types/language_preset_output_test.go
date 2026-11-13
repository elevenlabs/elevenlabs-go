package types

import (
	"testing"
)

func TestLanguagePresetOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "overrides": {
    "asr": {
      "keywords": [
        "value"
      ]
    },
    "turn": {
      "soft_timeout_config": {
        "message": "value",
        "additional_soft_timeout_messages": [
          "value"
        ]
      },
      "noisy_environment_message": "value"
    },
    "tts": {
      "model_id": "eleven_turbo_v2",
      "voice_id": "value",
      "supported_voices": [
        {
          "label": "value",
          "voice_id": "value",
          "description": "value",
          "language": "value",
          "model_family": "turbo",
          "optimize_streaming_latency": 123,
          "stability": 0.5,
          "speed": 1.2,
          "similarity_boost": 0.5,
          "clone_strength": 0.5
        }
      ],
      "stability": 9.99,
      "speed": 9.99,
      "similarity_boost": 9.99,
      "clone_strength": 9.99,
      "pronunciation_dictionary_locators": [
        {
          "pronunciation_dictionary_id": "value",
          "version_id": "value"
        }
      ]
    },
    "conversation": {
      "text_only": true,
      "max_duration_seconds": 123
    },
    "agent": {
      "first_message": "value",
      "language": "value",
      "max_conversation_duration_message": "value",
      "prompt": {
        "prompt": "value",
        "llm": "gpt-4o-mini",
        "reasoning_effort": "none",
        "tool_ids": [
          "value"
        ],
        "tool_ids_strict": true,
        "native_mcp_server_ids": [
          "value"
        ],
        "knowledge_base": [
          {
            "type": "file",
            "name": "value",
            "id": "value",
            "usage_mode": "prompt"
          }
        ]
      }
    }
  },
  "first_message_translation": {
    "source_hash": "value",
    "text": "value"
  },
  "soft_timeout_translation": {
    "source_hash": "value",
    "text": "value"
  },
  "noisy_environment_message_translation": {
    "source_hash": "value",
    "text": "value"
  }
}`
	assertRoundTrip(t, wire, new(LanguagePresetOutput))
}
