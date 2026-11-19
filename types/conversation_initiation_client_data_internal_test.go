package types

import (
	"testing"
)

func TestConversationInitiationClientDataInternalRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_config_override": {
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
  "custom_llm_extra_body": {
    "key": {}
  },
  "user_id": "value",
  "source_info": {
    "source": "unknown",
    "version": "value"
  },
  "branch_id": "value",
  "environment": "value",
  "starting_workflow_node_id": "value",
  "procedure_ids": [
    "value"
  ],
  "start_procedure_ids": [
    "value"
  ],
  "dynamic_variables": {
    "key": {}
  },
  "tool_mock_config": {
    "mocking_strategy": "all",
    "fallback_strategy": "call_real_tool",
    "mocked_tool_names": [
      "value"
    ]
  },
  "tool_mock_overrides": {
    "key": [
      {
        "parameter_conditions": [
          {
            "eval": {
              "type": "llm",
              "description": "value"
            },
            "path": "value"
          }
        ],
        "mock_result": "value",
        "is_error": true
      }
    ]
  }
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataInternal))
}
