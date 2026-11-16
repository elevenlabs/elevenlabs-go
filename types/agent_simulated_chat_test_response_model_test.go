package types

import (
	"testing"
)

func TestAgentSimulatedChatTestResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "simulated_conversation": [
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
      ],
      "file_input": {
        "file_id": "value",
        "original_filename": "value",
        "mime_type": "value",
        "file_url": "value"
      },
      "file_inputs": [
        {
          "file_id": "value",
          "original_filename": "value",
          "mime_type": "value",
          "file_url": "value"
        }
      ],
      "contextual_update_info": {
        "context_id": "value",
        "is_superseded": true
      },
      "reasoned": true
    }
  ],
  "analysis": {
    "evaluation_criteria_results": {
      "key": {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    },
    "data_collection_results": {
      "key": {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
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
        },
        "rationale": "value"
      }
    },
    "evaluation_criteria_results_list": [
      {
        "criteria_id": "value",
        "result": "success",
        "rationale": "value",
        "scoring_mode": "binary",
        "score": 123,
        "max_score": 123
      }
    ],
    "data_collection_results_list": [
      {
        "data_collection_id": "value",
        "name": "value",
        "value": {},
        "json_schema": {
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
        },
        "rationale": "value"
      }
    ],
    "call_successful": "success",
    "call_success_score": 9.99,
    "transcript_summary": "value",
    "call_summary_title": "value",
    "scoped": [
      {
        "scope": "conversation",
        "source_agent_id": "value",
        "source_branch_id": "value",
        "evaluation_criteria_results": {
          "key": {
            "criteria_id": "value",
            "result": "success",
            "rationale": "value",
            "scoring_mode": "binary",
            "score": 123,
            "max_score": 123
          }
        },
        "data_collection_results": {
          "key": {
            "data_collection_id": "value",
            "name": "value",
            "value": {},
            "json_schema": {
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
            },
            "rationale": "value"
          }
        },
        "successful": "success",
        "success_score": 9.99
      }
    ]
  }
}`
	assertRoundTrip(t, wire, new(AgentSimulatedChatTestResponseModel))
}
