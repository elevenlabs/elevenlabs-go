package agents

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCreateTestsRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  ],
  "name": "value",
  "parent_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateTestsRequestBody))
}

func TestBodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "entity_ids": [
    "value"
  ],
  "move_to": "value"
}`
	assertRoundTrip(t, wire, new(BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost))
}

func TestGetTestsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
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
  ],
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(GetTestsResponse))
}

func TestUpdateTestsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
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
  ],
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(UpdateTestsResponse))
}

func TestUpdateTestsRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  ],
  "name": "value",
  "parent_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(UpdateTestsRequestBody))
}

func TestListTestsByIdsRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "test_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ListTestsByIdsRequestModel))
}

func TestListTestsRequestSortModeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListTestsRequestSortMode
		wire    string
	}
	cases := []enumCase{
		{variant: ListTestsRequestSortModeDefault, wire: "\"default\""},
		{variant: ListTestsRequestSortModeFoldersFirst, wire: "\"folders_first\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListTestsRequestSortMode
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
