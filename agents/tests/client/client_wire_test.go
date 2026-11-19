package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{
  "id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateTestsRequestBody
	if err := json.Unmarshal([]byte(`{
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{
  "id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value"
}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agent-testing/create" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/create")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request agents.CreateTestsRequestBody
	if err := json.Unmarshal([]byte(`{
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestMoveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost
	if err := json.Unmarshal([]byte(`{
  "entity_ids": [
    "value"
  ],
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Move(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agent-testing/bulk-move" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/bulk-move")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestMoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request agents.BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost
	if err := json.Unmarshal([]byte(`{
  "entity_ids": [
    "value"
  ],
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Move(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agent-testing/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Get(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdateTestsRequestBody
	if err := json.Unmarshal([]byte(`{
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/convai/agent-testing/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request agents.UpdateTestsRequestBody
	if err := json.Unmarshal([]byte(`{
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/agent-testing/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Delete(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestSummariesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{
  "tests": {
    "key": {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.ListTestsByIdsRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Summaries(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{
  "tests": {
    "key": {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tests": {
    "key": {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  }
}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/agent-testing/summaries" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing/summaries")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSummariesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request agents.ListTestsByIdsRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Summaries(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsClient(t, 200, `{
  "tests": [
    {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	searchValue := string("value")
	parentFolderIdValue := string("value")
	includeFoldersValue := bool(true)
	createdByUserIdValue := string("value")
	request := &agents.ListTestsRequest{
		Cursor:          &cursorValue,
		PageSize:        &pageSizeValue,
		Search:          &searchValue,
		ParentFolderId:  &parentFolderIdValue,
		IncludeFolders:  &includeFoldersValue,
		CreatedByUserId: &createdByUserIdValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsClient(gotResponse, []byte(`{
  "tests": [
    {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tests": [
    {
      "id": "value",
      "name": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "created_at_unix_secs": 123,
      "last_updated_at_unix_secs": 123,
      "type": "llm",
      "entity_type": "test",
      "folder_parent_id": "value",
      "folder_path": [
        {
          "id": "value",
          "name": "value"
        }
      ],
      "children_count": 123,
      "conversation_initiation_source": "unknown"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsTestsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agent-testing" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent-testing")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("parent_folder_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "parent_folder_id", got, "value")
	}
	if got := req.Query.Get("include_folders"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_folders", got, "true")
	}
	if got := req.Query.Get("created_by_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "created_by_user_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	searchValue := string("value")
	parentFolderIdValue := string("value")
	includeFoldersValue := bool(true)
	createdByUserIdValue := string("value")
	request := &agents.ListTestsRequest{
		Cursor:          &cursorValue,
		PageSize:        &pageSizeValue,
		Search:          &searchValue,
		ParentFolderId:  &parentFolderIdValue,
		IncludeFolders:  &includeFoldersValue,
		CreatedByUserId: &createdByUserIdValue,
	}
	_, err := client.List(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

type wireRequestAgentsTestsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsTestsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsTestsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsTestsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsTestsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsTestsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if responseBody != "" {
			_, _ = w.Write([]byte(responseBody))
		}
	}))
	t.Cleanup(server.Close)
	return server, recorded
}

func receiveWireAgentsTestsClient(t *testing.T, recorded <-chan wireRequestAgentsTestsClient) wireRequestAgentsTestsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsTestsClient{}
	}
}

// wireJSONContainsAgentsTestsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsTestsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsTestsClient(got, want)
	}
	for key, wantValue := range wantObject {
		if key == "stream" {
			continue
		}
		if gotValue, ok := gotObject[key]; !ok || !reflect.DeepEqual(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// setWireValueAgentsTestsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsTestsClient(t *testing.T, target any, value any) {
	t.Helper()
	field := reflect.ValueOf(target).Elem()
	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}
	switch typed := value.(type) {
	case string:
		field.SetString(typed)
	case int:
		field.SetInt(int64(typed))
	case float64:
		field.SetFloat(typed)
	default:
		t.Fatalf("unsupported value %T", value)
	}
}

func wireJSONBytesEqualAgentsTestsClient(got, want []byte) bool {
	var gotValue any
	var wantValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		return false
	}
	if err := json.Unmarshal(want, &wantValue); err != nil {
		return false
	}
	return reflect.DeepEqual(gotValue, wantValue)
}
