package types

import (
	"testing"
)

func TestBodySimulatesAConversationV1ConvaiAgentsAgentIdSimulateConversationPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "simulation_specification": {
    "simulated_user_config": {
      "first_message": "value",
      "language": "value",
      "hinglish_mode": true,
      "dynamic_variables": {},
      "disable_first_message_interruptions": true,
      "max_conversation_duration_message": "value",
      "text_behavior_overrides": {
        "key": {
          "verbosity": "auto",
          "output_format": "mp3_22050_32",
          "interaction_budget": "realtime"
        }
      },
      "subagent_config": {
        "agents": [
          {
            "execution_mode": "immediate",
            "timeout_secs": 123,
            "agent_id": "value",
            "branch_id": "value",
            "description": "value",
            "parameters": {
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
            }
          }
        ],
        "self_subagent": {
          "execution_mode": "immediate",
          "timeout_secs": 123,
          "description": "value"
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
          "max_retrieved_rag_chunks_count": 20,
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
    },
    "tool_mock_config": {
      "key": {
        "default_return_value": "value",
        "default_is_error": true
      }
    },
    "partial_conversation_history": [
      {
        "role": "user",
        "agent_metadata": {
          "agent_id": "value",
          "branch_id": "value",
          "workflow_node_id": "value",
          "version_id": "value"
        },
        "message": "value",
        "multivoice_message": {
          "parts": [
            {
              "text": "value",
              "voice_label": "value",
              "time_in_call_secs": 123
            }
          ]
        },
        "tool_calls": [
          {
            "type": "system",
            "request_id": "value",
            "tool_name": "value",
            "tool_id": "value",
            "params_as_json": "value",
            "tool_has_been_called": true,
            "tool_details": {
              "type": "webhook",
              "method": "value",
              "url": "value",
              "headers": {
                "key": "value"
              },
              "path_params": {
                "key": "value"
              },
              "query_params": {
                "key": "value"
              },
              "body": "value"
            }
          }
        ],
        "tool_results": [
          {
            "request_id": "value",
            "tool_name": "value",
            "tool_id": "value",
            "result_value": "value",
            "is_error": true,
            "is_blocked": true,
            "tool_has_been_called": true,
            "tool_latency_secs": 9.99,
            "error_type": "value",
            "raw_error_message": "value",
            "dynamic_variable_updates": [
              {
                "variable_name": "value",
                "old_value": "value",
                "new_value": "value",
                "updated_at": 9.99,
                "tool_name": "value",
                "tool_request_id": "value"
              }
            ],
            "type": "client"
          }
        ],
        "feedback": {
          "score": "like",
          "time_in_call_secs": 123
        },
        "llm_override": "value",
        "producing_llm": "value",
        "time_in_call_secs": 123,
        "conversation_turn_metrics": {
          "metrics": {
            "key": {
              "elapsed_time": 9.99
            }
          },
          "convai_asr_provider": "value",
          "convai_tts_model": "value",
          "convai_tts_cascade": "value",
          "voice_id": "value"
        },
        "rag_retrieval_info": {
          "chunks": [
            {
              "document_id": "value",
              "chunk_id": "value",
              "vector_distance": 9.99
            }
          ],
          "embedding_model": "e5_mistral_7b_instruct",
          "retrieval_query": "value",
          "rag_latency_secs": 9.99,
          "used_chunk_ids": [
            "value"
          ]
        },
        "llm_usage": {
          "model_usage": {
            "key": {
              "input": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_read": {
                "tokens": 123,
                "price": 9.99
              },
              "input_cache_write": {
                "tokens": 123,
                "price": 9.99
              },
              "output_total": {
                "tokens": 123,
                "price": 9.99
              }
            }
          }
        },
        "interrupted": true,
        "ignored_as_backchannel": true,
        "platform_event": "user_turn_timeout",
        "original_message": "value",
        "reasoning": [
          {
            "summary": "value",
            "provider_redact": true
          }
        ],
        "source_medium": "audio",
        "source_event_id": 123,
        "used_static_kb_document_ids": [
          "value"
        ],
        "user_identifier": "value",
        "id": "value",
        "created_at_unix_ms": 123,
        "triggered_guardrails": [
          {
            "guardrail_type": "custom",
            "guardrail_name": "value",
            "execution_mode": "streaming",
            "guardrail_details": {
              "type": "llm",
              "rationale": "value"
            }
          }
        ]
      }
    ],
    "dynamic_variables": {
      "key": {}
    }
  },
  "extra_evaluation_criteria": [
    {
      "id": "value",
      "name": "value",
      "type": "prompt",
      "conversation_goal_prompt": "value",
      "use_knowledge_base": true,
      "scope": "conversation",
      "llm": "gpt-4o-mini",
      "scoring_mode": "binary",
      "max_score": 123,
      "score_instructions": "value"
    }
  ],
  "new_turns_limit": 100
}`
	assertRoundTrip(t, wire, new(BodySimulatesAConversationV1ConvaiAgentsAgentIdSimulateConversationPost))
}
