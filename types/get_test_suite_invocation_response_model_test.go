package types

import (
	"testing"
)

func TestGetTestSuiteInvocationResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "ran_against_draft": true,
  "created_at": 123,
  "folder_id": "value",
  "repeat_count": 123,
  "cancelled": true,
  "bucketing_status": "pending",
  "result_groups": [
    {
      "test_id": "value",
      "test_name": "value",
      "workflow_node_id": "value",
      "buckets": [
        {
          "test_run_ids": [
            "value"
          ],
          "title": "value",
          "reason": "value",
          "status": "pending"
        }
      ]
    }
  ],
  "test_runs": [
    {
      "test_run_id": "value",
      "test_info": {
        "from_conversation_metadata": {
          "conversation_id": "value",
          "agent_id": "value",
          "branch_id": "value",
          "workflow_node_id": "value",
          "original_agent_reply": [
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
          ]
        },
        "dynamic_variables": {
          "key": {}
        },
        "chat_history": [
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
        "conversation_initiation_source": "unknown",
        "environment": "value",
        "type": "llm",
        "success_condition": "value",
        "success_examples": [
          {
            "response": "value",
            "type": "success"
          }
        ],
        "failure_examples": [
          {
            "response": "value",
            "type": "failure"
          }
        ]
      },
      "test_invocation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "workflow_node_id": "value",
      "status": "pending",
      "agent_responses": [
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
      ],
      "test_id": "value",
      "test_name": "value",
      "condition_result": {
        "result": "success",
        "rationale": {
          "messages": [
            "value"
          ],
          "summary": "value"
        }
      },
      "last_updated_at_unix": 123,
      "metadata": {
        "workspace_id": "value",
        "test_name": "value",
        "ran_by_user_email": "value",
        "test_type": "llm"
      },
      "root_folder_id": "value",
      "root_folder_name": "value",
      "environment": "value",
      "credits_used": 123,
      "charging": {
        "dev_discount": true,
        "is_burst": true,
        "tier": "value",
        "llm_usage": {
          "irreversible_generation": {
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
          "initiated_generation": {
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
          }
        },
        "llm_price": 9.99,
        "llm_charge": 123,
        "call_charge": 123,
        "platform_charge": 123,
        "platform_usage": {
          "category_usage": {
            "key": {
              "credits": 123,
              "price": 9.99,
              "quantity": 9.99
            }
          }
        },
        "platform_price": 9.99,
        "free_minutes_consumed": 9.99,
        "free_llm_dollars_consumed": 9.99,
        "tts_usage": {
          "primary_tts_model": "value",
          "total_audio_output_seconds": 9.99,
          "total_characters": 123,
          "per_voice_usage": [
            {
              "voice_id": "value",
              "audio_output_seconds": 9.99
            }
          ]
        },
        "asr_usage": {
          "asr_model": "value",
          "total_transcription_calls": 123,
          "total_audio_input_seconds": 9.99
        },
        "analysis": {
          "total": {
            "price": 9.99,
            "charge": 123,
            "runs": 123,
            "price_per_feature": {
              "key": 9.99
            },
            "charge_per_feature": {
              "key": 123
            }
          },
          "last_run": {
            "price": 9.99,
            "charge": 123,
            "price_per_feature": {
              "key": 9.99
            },
            "charge_per_feature": {
              "key": 123
            }
          }
        },
        "voice_surcharge_charge": 123
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetTestSuiteInvocationResponseModel))
}

func TestGetTestSuiteInvocationResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "ran_against_draft": true,
  "created_at": 123,
  "folder_id": "value",
  "repeat_count": 123,
  "cancelled": true,
  "bucketing_status": "pending",
  "result_groups": [
    {
      "test_id": "value",
      "test_name": "value",
      "workflow_node_id": "value",
      "buckets": [
        {
          "test_run_ids": [
            "value"
          ],
          "title": "value",
          "reason": "value",
          "status": "pending"
        }
      ]
    }
  ],
  "test_runs": [
    {
      "test_run_id": "value",
      "test_info": {
        "from_conversation_metadata": {
          "conversation_id": "value",
          "agent_id": "value",
          "branch_id": "value",
          "workflow_node_id": "value",
          "original_agent_reply": [
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
          ]
        },
        "dynamic_variables": {
          "key": {}
        },
        "chat_history": [
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
        "conversation_initiation_source": "unknown",
        "environment": "value",
        "type": "llm",
        "success_condition": "value",
        "success_examples": [
          {
            "response": "value",
            "type": "success"
          }
        ],
        "failure_examples": [
          {
            "response": "value",
            "type": "failure"
          }
        ]
      },
      "test_invocation_id": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "workflow_node_id": "value",
      "status": "pending",
      "agent_responses": [
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
      ],
      "test_id": "value",
      "test_name": "value",
      "condition_result": {
        "result": "success",
        "rationale": {
          "messages": [
            "value"
          ],
          "summary": "value"
        }
      },
      "last_updated_at_unix": 123,
      "metadata": {
        "workspace_id": "value",
        "test_name": "value",
        "ran_by_user_email": "value",
        "test_type": "llm"
      },
      "root_folder_id": "value",
      "root_folder_name": "value",
      "environment": "value",
      "credits_used": 123,
      "charging": {
        "dev_discount": true,
        "is_burst": true,
        "tier": "value",
        "llm_usage": {
          "irreversible_generation": {
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
          "initiated_generation": {
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
          }
        },
        "llm_price": 9.99,
        "llm_charge": 123,
        "call_charge": 123,
        "platform_charge": 123,
        "platform_usage": {
          "category_usage": {
            "key": {
              "credits": 123,
              "price": 9.99,
              "quantity": 9.99
            }
          }
        },
        "platform_price": 9.99,
        "free_minutes_consumed": 9.99,
        "free_llm_dollars_consumed": 9.99,
        "tts_usage": {
          "primary_tts_model": "value",
          "total_audio_output_seconds": 9.99,
          "total_characters": 123,
          "per_voice_usage": [
            {
              "voice_id": "value",
              "audio_output_seconds": 9.99
            }
          ]
        },
        "asr_usage": {
          "asr_model": "value",
          "total_transcription_calls": 123,
          "total_audio_input_seconds": 9.99
        },
        "analysis": {
          "total": {
            "price": 9.99,
            "charge": 123,
            "runs": 123,
            "price_per_feature": {
              "key": 9.99
            },
            "charge_per_feature": {
              "key": 123
            }
          },
          "last_run": {
            "price": 9.99,
            "charge": 123,
            "price_per_feature": {
              "key": 9.99
            },
            "charge_per_feature": {
              "key": 123
            }
          }
        },
        "voice_surcharge_charge": 123
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetTestSuiteInvocationResponseModel))
}
