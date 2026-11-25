package types

import (
	"testing"
)

func TestPromptAgentApiModelWorkflowOverrideInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt": "value",
  "llm": "gpt-4o-mini",
  "reasoning_effort": "none",
  "service_tier": "standard",
  "thinking_budget": 123,
  "enable_reasoning_summary": true,
  "temperature": 9.99,
  "max_tokens": 123,
  "tool_ids": [
    "value"
  ],
  "tool_ids_strict": true,
  "built_in_tools": {
    "transfer_to_agent": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "end_call": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "language_detection": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "transfer_to_number": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "skip_turn": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "play_keypad_touch_tone": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    },
    "voicemail_detection": {
      "type": "system",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "params": {
        "system_tool_type": "end_call"
      }
    }
  },
  "enable_parallel_tool_calls": true,
  "mcp_server_ids": [
    "value"
  ],
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
  ],
  "custom_llm": {
    "url": "value",
    "model_id": "value",
    "api_key": {
      "secret_id": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "mtls_auth_connection": {
      "auth_connection_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "api_version": "value",
    "api_type": "chat_completions"
  },
  "ignore_default_personality": true,
  "rag": {
    "enabled": true,
    "embedding_model": "e5_mistral_7b_instruct",
    "max_vector_distance": 9.99,
    "max_documents_length": 123,
    "max_retrieved_rag_chunks_count": 123,
    "num_candidates": 123,
    "query_rewrite_prompt_override": "value",
    "knowledge_base_tool_info": {
      "enabled_strategies": [
        "cat"
      ]
    }
  },
  "timezone": "value",
  "backup_llm_config": {
    "preference": "default"
  },
  "cascade_timeout_seconds": 9.99,
  "tools": [
    {
      "type": "webhook",
      "name": "value",
      "description": "value",
      "response_timeout_secs": 123,
      "disable_interruptions": true,
      "interruption_mode": "allow",
      "force_pre_tool_speech": true,
      "pre_tool_speech": "auto",
      "assignments": [
        {
          "source": "response",
          "dynamic_variable": "value",
          "value_path": "value",
          "sanitize": true,
          "preserve_native_type": true
        }
      ],
      "tool_call_sound": "typing",
      "tool_call_sound_behavior": "auto",
      "tool_error_handling_mode": "auto",
      "dynamic_variables": {
        "dynamic_variable_placeholders": {
          "key": {}
        }
      },
      "execution_mode": "immediate",
      "api_schema": {
        "request_headers": {
          "key": "value"
        },
        "url": "value",
        "method": "GET",
        "path_params_schema": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "query_params_schema": {
          "properties": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "required": [
            "value"
          ]
        },
        "request_body_schema": {
          "property_kind": "array",
          "description": "value",
          "dynamic_variable": "value",
          "constant_value": {
            "key": {}
          },
          "is_omitted": true,
          "type": "object",
          "required": [
            "value"
          ],
          "properties": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "required_constraints": {
            "any_of": [
              {
                "required": [
                  "value"
                ]
              }
            ],
            "all_of": [
              {
                "required": [
                  "value"
                ]
              }
            ]
          }
        },
        "response_body_schema": {
          "property_kind": "array",
          "description": "value",
          "dynamic_variable": "value",
          "constant_value": {
            "key": {}
          },
          "is_omitted": true,
          "type": "object",
          "required": [
            "value"
          ],
          "properties": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "required_constraints": {
            "any_of": [
              {
                "required": [
                  "value"
                ]
              }
            ],
            "all_of": [
              {
                "required": [
                  "value"
                ]
              }
            ]
          }
        },
        "response_filter": {
          "mode": "all",
          "filters": [
            "value"
          ],
          "content_type": "application/json"
        },
        "content_type": "application/json",
        "auth_resolved_params": [
          "value"
        ],
        "auth_connection": {
          "auth_connection_id": "value"
        }
      },
      "follow_redirects": true,
      "follow_redirects_allowed_domains": [
        "value"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelWorkflowOverrideInput))
}
