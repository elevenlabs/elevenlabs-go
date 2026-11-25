package types

import (
	"testing"
)

func TestWorkflowOverrideAgentNodeModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "override_agent",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "conversation_config": {
    "asr": {
      "quality": "high",
      "provider": "elevenlabs",
      "user_input_audio_format": "pcm_8000",
      "keywords": [
        "value"
      ]
    },
    "turn": {
      "turn_timeout": 9.99,
      "initial_wait_time": 9.99,
      "silence_end_call_timeout": 9.99,
      "turn_eagerness": "patient",
      "spelling_patience": "auto",
      "speculative_turn": true,
      "retranscribe_on_turn_timeout": true,
      "turn_model": "turn_v2",
      "interruption_ignore_terms": [
        "value"
      ],
      "interruption_ignore_term_languages": [
        "value"
      ],
      "merge_with_default_ignore_terms": true,
      "transcribe_on_disabled_interruptions": true,
      "soft_timeout_config": {
        "timeout_seconds": 9.99,
        "message": "value",
        "additional_soft_timeout_messages": [
          "value"
        ],
        "use_llm_generated_message": true,
        "randomize_fillers": true,
        "max_soft_timeouts_per_generation": 123,
        "pause_after_filler": true,
        "llm_generated_message_prompt_override": "value",
        "disable_until_first_user_message": true
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
      "expressive_mode": true,
      "suggested_audio_tags": [
        {
          "tag": "value",
          "description": "value"
        }
      ],
      "agent_output_audio_format": "pcm_8000",
      "optimize_streaming_latency": 123,
      "stability": 9.99,
      "speed": 9.99,
      "similarity_boost": 9.99,
      "clone_strength": 9.99,
      "text_normalisation_type": "system_prompt",
      "pronunciation_dictionary_locators": [
        {
          "pronunciation_dictionary_id": "value",
          "version_id": "value"
        }
      ],
      "enable_phoneme_tags": true,
      "audio_effects": {
        "filter_preset_id": "value",
        "distance": 0.5,
        "environment_id": "value",
        "background_noise_id": "value",
        "send_level": 0.5,
        "pan": 0.5
      }
    },
    "conversation": {
      "text_only": true,
      "max_duration_seconds": 123,
      "client_events": [
        "conversation_initiation_metadata"
      ],
      "file_input": {
        "enabled": true,
        "max_files_in_memory": 123,
        "max_files_per_conversation": 123
      },
      "monitoring_enabled": true,
      "monitoring_events": [
        "conversation_initiation_metadata"
      ],
      "dtmf_input_settings": {
        "dtmf_input_timeout": 9.99,
        "hash_terminator": true,
        "redact_input": true
      },
      "background_sound": {
        "source_type": "preset",
        "source_id": "office2",
        "volume": 9.99,
        "crossfade_loop": true
      },
      "source_attribution": true
    },
    "language_presets": {
      "key": {
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
      }
    },
    "vad": {},
    "agent": {
      "first_message": "value",
      "language": "value",
      "hinglish_mode": true,
      "dynamic_variables": {
        "dynamic_variable_placeholders": {
          "key": {}
        }
      },
      "disable_first_message_interruptions": true,
      "max_conversation_duration_message": "value",
      "text_behavior_overrides": {
        "key": {
          "verbosity": "auto",
          "output_format": "mp3_22050_32",
          "interaction_budget": "realtime"
        }
      },
      "prompt": {
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
      }
    }
  },
  "additional_prompt": "value",
  "additional_knowledge_base": [
    {
      "type": "file",
      "name": "value",
      "id": "value",
      "usage_mode": "prompt"
    }
  ],
  "additional_tool_ids": [
    "value"
  ],
  "label": "value",
  "entry_behavior": "generate_immediately",
  "allow_interruptions_on_entry": true
}`
	assertRoundTrip(t, wire, new(WorkflowOverrideAgentNodeModelInput))
}
