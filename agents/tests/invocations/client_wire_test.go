package invocations

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

	tests "github.com/elevenlabs/elevenlabs-go/agents/tests"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsInvocations(t, 200, `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "id": "value",
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "runs_diverged_from_version": true,
      "created_at_unix_secs": 123,
      "test_run_count": 123,
      "passed_count": 123,
      "failed_count": 123,
      "pending_count": 123,
      "cancelled_count": 123,
      "title": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "repeat_count": 123,
      "cancelled": true,
      "credits_used": 123,
      "total_price": 9.99
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	branchIdValue := string("value")
	createdByUserIdValue := string("value")
	testIdValue := string("value")
	startTimeValue := int(123)
	endTimeValue := int(123)
	versionIdValue := string("value")
	pageSizeValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &tests.ListInvocationsRequest{
		AgentId:         &agentIdValue,
		BranchId:        &branchIdValue,
		CreatedByUserId: &createdByUserIdValue,
		TestId:          &testIdValue,
		StartTime:       &startTimeValue,
		EndTime:         &endTimeValue,
		VersionId:       &versionIdValue,
		PageSize:        &pageSizeValue,
		Search:          &searchValue,
		Cursor:          &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsInvocations(gotResponse, []byte(`{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "id": "value",
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "runs_diverged_from_version": true,
      "created_at_unix_secs": 123,
      "test_run_count": 123,
      "passed_count": 123,
      "failed_count": 123,
      "pending_count": 123,
      "cancelled_count": 123,
      "title": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "repeat_count": 123,
      "cancelled": true,
      "credits_used": 123,
      "total_price": 9.99
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
      "id": "value",
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "version_id": "value",
      "ran_against_draft": true,
      "runs_diverged_from_version": true,
      "created_at_unix_secs": 123,
      "test_run_count": 123,
      "passed_count": 123,
      "failed_count": 123,
      "pending_count": 123,
      "cancelled_count": 123,
      "title": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "repeat_count": 123,
      "cancelled": true,
      "credits_used": 123,
      "total_price": 9.99
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsTestsInvocations(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/test-invocations" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/test-invocations")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("created_by_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "created_by_user_id", got, "value")
	}
	if got := req.Query.Get("test_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "test_id", got, "value")
	}
	if got := req.Query.Get("start_time"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "start_time", got, "123")
	}
	if got := req.Query.Get("end_time"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "end_time", got, "123")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsInvocations(t, 422, `{
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

	agentIdValue := string("value")
	branchIdValue := string("value")
	createdByUserIdValue := string("value")
	testIdValue := string("value")
	startTimeValue := int(123)
	endTimeValue := int(123)
	versionIdValue := string("value")
	pageSizeValue := int(100)
	searchValue := string("value")
	cursorValue := string("value")
	request := &tests.ListInvocationsRequest{
		AgentId:         &agentIdValue,
		BranchId:        &branchIdValue,
		CreatedByUserId: &createdByUserIdValue,
		TestId:          &testIdValue,
		StartTime:       &startTimeValue,
		EndTime:         &endTimeValue,
		VersionId:       &versionIdValue,
		PageSize:        &pageSizeValue,
		Search:          &searchValue,
		Cursor:          &cursorValue,
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsInvocations(t, 200, `{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
	if !wireJSONBytesEqualAgentsTestsInvocations(gotResponse, []byte(`{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
}`)
	}

	req := receiveWireAgentsTestsInvocations(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/test-invocations/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/test-invocations/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsInvocations(t, 422, `{
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

	server, recorded := serveWireAgentsTestsInvocations(t, 200, `{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request tests.UpdateTestSuiteInvocationRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_invocation_id": "value",
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsInvocations(gotResponse, []byte(`{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
                    },
                    "output_total_reasoning_portion": {
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
                  },
                  "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
                },
                "output_total_reasoning_portion": {
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
}`)
	}

	req := receiveWireAgentsTestsInvocations(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/test-invocations/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/test-invocations/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsInvocations(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsInvocations(t, 422, `{
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

	var request tests.UpdateTestSuiteInvocationRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_invocation_id": "value",
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
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

func TestCancelWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsInvocations(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Cancel(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsTestsInvocations(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/test-invocations/value/cancel" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/test-invocations/value/cancel")
	}
}

func TestCancelWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsInvocations(t, 422, `{
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

	err := client.Cancel(context.Background(), string("value"))
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

func TestResubmitWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsTestsInvocations(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request tests.ResubmitTestsRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_invocation_id": "value",
  "test_run_ids": [
    "value"
  ],
  "agent_config_override": {
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
          "max_soft_timeouts_per_generation": 8,
          "pause_after_filler": true,
          "llm_generated_message_prompt_override": "value",
          "disable_until_first_user_message": true
        },
        "noisy_environment_detection": true,
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
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5,
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
          "max_files_in_memory": 30,
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
          "volume": 0.5,
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
      }
    },
    "platform_settings": {
      "evaluation": {
        "criteria": [
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
        ]
      },
      "widget": {
        "variant": "tiny",
        "placement": "top-left",
        "expandable": "never",
        "avatar": {
          "type": "orb",
          "color_1": "value",
          "color_2": "value"
        },
        "feedback_mode": "none",
        "end_feedback": {
          "type": "rating"
        },
        "bg_color": "value",
        "text_color": "value",
        "btn_color": "value",
        "btn_text_color": "value",
        "border_color": "value",
        "focus_color": "value",
        "border_radius": 123,
        "btn_radius": 123,
        "action_text": "value",
        "start_call_text": "value",
        "end_call_text": "value",
        "expand_text": "value",
        "listening_text": "value",
        "speaking_text": "value",
        "shareable_page_text": "value",
        "shareable_page_show_terms": true,
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value",
        "show_avatar_when_collapsed": true,
        "disable_banner": true,
        "override_link": "value",
        "markdown_link_allowed_hosts": [
          {
            "hostname": "value"
          }
        ],
        "markdown_link_include_www": true,
        "markdown_link_allow_http": true,
        "mic_muting_enabled": true,
        "transcript_enabled": true,
        "text_input_enabled": true,
        "conversation_mode_toggle_enabled": true,
        "default_expanded": true,
        "always_expanded": true,
        "dismissible": true,
        "show_agent_status": true,
        "show_conversation_id": true,
        "strip_audio_tags": true,
        "syntax_highlight_theme": "light",
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "styles": {
          "base": "value",
          "base_hover": "value",
          "base_active": "value",
          "base_border": "value",
          "base_subtle": "value",
          "base_primary": "value",
          "base_error": "value",
          "accent": "value",
          "accent_hover": "value",
          "accent_active": "value",
          "accent_border": "value",
          "accent_subtle": "value",
          "accent_primary": "value",
          "overlay_padding": 9.99,
          "button_radius": 9.99,
          "input_radius": 9.99,
          "bubble_radius": 9.99,
          "sheet_radius": 9.99,
          "compact_sheet_radius": 9.99,
          "dropdown_sheet_radius": 9.99
        },
        "show_resize_button": true,
        "end_confirmation_enabled": true,
        "mobile_fullscreen": true,
        "transcript_speech_sync_enabled": true,
        "language_selector": true,
        "supports_text_only": true,
        "custom_avatar_path": "value",
        "language_presets": {
          "key": {
            "text_contents": {
              "main_label": "value",
              "start_call": "value",
              "start_chat": "value",
              "new_call": "value",
              "end_call": "value",
              "mute_microphone": "value",
              "change_language": "value",
              "collapse": "value",
              "expand": "value",
              "copied": "value",
              "accept_terms": "value",
              "dismiss_terms": "value",
              "listening_status": "value",
              "speaking_status": "value",
              "connecting_status": "value",
              "chatting_status": "value",
              "queue_waiting_status": "value",
              "queue_waiting_status_short": "value",
              "input_label": "value",
              "input_placeholder": "value",
              "input_placeholder_text_only": "value",
              "input_placeholder_new_conversation": "value",
              "user_ended_conversation": "value",
              "agent_ended_conversation": "value",
              "conversation_id": "value",
              "error_occurred": "value",
              "queue_timed_out": "value",
              "copy_id": "value",
              "initiate_feedback": "value",
              "request_follow_up_feedback": "value",
              "thanks_for_feedback": "value",
              "thanks_for_feedback_details": "value",
              "follow_up_feedback_placeholder": "value",
              "submit": "value",
              "go_back": "value",
              "send_message": "value",
              "text_mode": "value",
              "voice_mode": "value",
              "switched_to_text_mode": "value",
              "switched_to_voice_mode": "value",
              "copy": "value",
              "download": "value",
              "wrap": "value",
              "agent_working": "value",
              "agent_done": "value",
              "agent_error": "value",
              "attach_file": "value",
              "remove_file": "value",
              "file_upload_error": "value",
              "file_type_unsupported": "value",
              "file_too_large": "value",
              "file_limit_reached": "value",
              "typing_indicator": "value",
              "rich_content_unavailable": "value"
            },
            "text_contents_translation": {
              "source": {
                "key": "value"
              },
              "text": {
                "key": "value"
              }
            },
            "terms_text": "value",
            "terms_html": "value",
            "terms_key": "value",
            "terms_translation": {
              "source_hash": "value",
              "text": "value"
            }
          }
        }
      },
      "data_collection": {
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
          "is_omitted": true,
          "name": "value",
          "llm": "gpt-4o-mini"
        }
      },
      "data_collection_scopes": {
        "key": "conversation"
      },
      "analysis_items": {
        "evaluation_criteria": [
          {
            "source": "user",
            "analysis_item_id": "value",
            "version_id": "value",
            "additional_version_ids": [
              "value"
            ],
            "scope": "conversation",
            "weight": 9.99
          }
        ],
        "data_collection": [
          {
            "source": "user",
            "analysis_item_id": "value",
            "version_id": "value",
            "scope": "conversation"
          }
        ]
      },
      "overrides": {
        "conversation_config_override": {
          "asr": {
            "keywords": true
          },
          "turn": {
            "soft_timeout_config": {
              "message": true,
              "additional_soft_timeout_messages": true
            },
            "noisy_environment_message": true
          },
          "tts": {
            "model_id": true,
            "voice_id": true,
            "supported_voices": true,
            "stability": true,
            "speed": true,
            "similarity_boost": true,
            "clone_strength": true,
            "pronunciation_dictionary_locators": true
          },
          "conversation": {
            "text_only": true,
            "max_duration_seconds": true
          },
          "agent": {
            "first_message": true,
            "language": true,
            "max_conversation_duration_message": true,
            "prompt": {
              "prompt": true,
              "llm": true,
              "reasoning_effort": true,
              "tool_ids": true,
              "tool_ids_strict": true,
              "native_mcp_server_ids": true,
              "knowledge_base": true
            }
          }
        },
        "custom_llm_extra_body": true,
        "enable_conversation_initiation_client_data_from_webhook": true,
        "enable_starting_workflow_node_id_from_client": true,
        "enable_procedure_ids_from_client": true
      },
      "workspace_overrides": {
        "conversation_initiation_client_data_webhook": {
          "url": "value",
          "request_headers": {
            "key": "value"
          }
        },
        "webhooks": {
          "post_call_webhook_id": "value",
          "events": [
            "transcript"
          ],
          "transcript_format": "json",
          "exclude_transcript": true,
          "send_audio": true
        }
      },
      "testing": {
        "attached_tests": [
          {
            "test_id": "value",
            "workflow_node_id": "value"
          }
        ]
      },
      "archived": true,
      "guardrails": {
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
      },
      "summary_language": "value",
      "auto_translate_transcript_to_app_language": true,
      "auth": {
        "enable_auth": true,
        "allowlist": [
          {
            "hostname": "value"
          }
        ],
        "require_origin_header": true,
        "shareable_token": "value"
      },
      "call_limits": {
        "agent_concurrency_limit": 123,
        "daily_limit": 123,
        "bursting_enabled": true
      },
      "queueing_config": {
        "enabled": true,
        "wait_timeout_seconds": 123,
        "hold_audio": {
          "audio_path": "value",
          "audio_url": "value",
          "original_filename": "value",
          "duration_secs": 9.99,
          "size_bytes": 123
        }
      },
      "privacy": {
        "record_voice": true,
        "retention_days": 123,
        "delete_transcript_and_pii": true,
        "delete_audio": true,
        "apply_to_existing_conversations": true,
        "zero_retention_mode": true,
        "conversation_history_redaction": {
          "enabled": true,
          "entities": [
            "name"
          ],
          "excluded_data_collection_ids": [
            "value"
          ]
        }
      },
      "trust_context": "unknown",
      "analysis_llm": "gpt-4o-mini",
      "topic_discovery": {},
      "sentiment_analysis": {},
      "alerting": {
        "monitor_configs": {
          "key": {
            "enabled": true,
            "threshold": 0.5,
            "relative_increase_threshold": 9.99,
            "min_failure_count": 123,
            "min_history_bucket_count": 123,
            "min_sample_count": 123,
            "suspect_trigger_threshold": 10,
            "auto_resolve_after_inactive_minutes": 123
          }
        },
        "auto_resolve_after_inactive_minutes": 123,
        "environments": [
          "value"
        ],
        "notifiers": [
          {
            "type": "webhook",
            "webhook_id": "value"
          }
        ]
      }
    },
    "workflow": {
      "edges": {
        "key": {
          "source": "value",
          "target": "value",
          "forward_condition": {
            "label": "value",
            "type": "unconditional"
          },
          "backward_condition": {
            "label": "value",
            "type": "unconditional"
          },
          "reroutes": [
            {
              "x": 9.99,
              "y": 9.99
            }
          ]
        }
      },
      "nodes": {
        "key": {
          "type": "start",
          "position": {
            "x": 9.99,
            "y": 9.99
          },
          "edge_order": [
            "value"
          ]
        }
      },
      "prevent_subagent_loops": true
    }
  },
  "agent_id": "value",
  "branch_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Resubmit(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsTestsInvocations(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsTestsInvocations(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/test-invocations/value/resubmit" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/test-invocations/value/resubmit")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsTestsInvocations(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestResubmitWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsTestsInvocations(t, 422, `{
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

	var request tests.ResubmitTestsRequestModel
	if err := json.Unmarshal([]byte(`{
  "test_invocation_id": "value",
  "test_run_ids": [
    "value"
  ],
  "agent_config_override": {
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
          "max_soft_timeouts_per_generation": 8,
          "pause_after_filler": true,
          "llm_generated_message_prompt_override": "value",
          "disable_until_first_user_message": true
        },
        "noisy_environment_detection": true,
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
        "stability": 0.5,
        "speed": 1.2,
        "similarity_boost": 0.5,
        "clone_strength": 0.5,
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
          "max_files_in_memory": 30,
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
          "volume": 0.5,
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
      }
    },
    "platform_settings": {
      "evaluation": {
        "criteria": [
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
        ]
      },
      "widget": {
        "variant": "tiny",
        "placement": "top-left",
        "expandable": "never",
        "avatar": {
          "type": "orb",
          "color_1": "value",
          "color_2": "value"
        },
        "feedback_mode": "none",
        "end_feedback": {
          "type": "rating"
        },
        "bg_color": "value",
        "text_color": "value",
        "btn_color": "value",
        "btn_text_color": "value",
        "border_color": "value",
        "focus_color": "value",
        "border_radius": 123,
        "btn_radius": 123,
        "action_text": "value",
        "start_call_text": "value",
        "end_call_text": "value",
        "expand_text": "value",
        "listening_text": "value",
        "speaking_text": "value",
        "shareable_page_text": "value",
        "shareable_page_show_terms": true,
        "terms_text": "value",
        "terms_html": "value",
        "terms_key": "value",
        "show_avatar_when_collapsed": true,
        "disable_banner": true,
        "override_link": "value",
        "markdown_link_allowed_hosts": [
          {
            "hostname": "value"
          }
        ],
        "markdown_link_include_www": true,
        "markdown_link_allow_http": true,
        "mic_muting_enabled": true,
        "transcript_enabled": true,
        "text_input_enabled": true,
        "conversation_mode_toggle_enabled": true,
        "default_expanded": true,
        "always_expanded": true,
        "dismissible": true,
        "show_agent_status": true,
        "show_conversation_id": true,
        "strip_audio_tags": true,
        "syntax_highlight_theme": "light",
        "text_contents": {
          "main_label": "value",
          "start_call": "value",
          "start_chat": "value",
          "new_call": "value",
          "end_call": "value",
          "mute_microphone": "value",
          "change_language": "value",
          "collapse": "value",
          "expand": "value",
          "copied": "value",
          "accept_terms": "value",
          "dismiss_terms": "value",
          "listening_status": "value",
          "speaking_status": "value",
          "connecting_status": "value",
          "chatting_status": "value",
          "queue_waiting_status": "value",
          "queue_waiting_status_short": "value",
          "input_label": "value",
          "input_placeholder": "value",
          "input_placeholder_text_only": "value",
          "input_placeholder_new_conversation": "value",
          "user_ended_conversation": "value",
          "agent_ended_conversation": "value",
          "conversation_id": "value",
          "error_occurred": "value",
          "queue_timed_out": "value",
          "copy_id": "value",
          "initiate_feedback": "value",
          "request_follow_up_feedback": "value",
          "thanks_for_feedback": "value",
          "thanks_for_feedback_details": "value",
          "follow_up_feedback_placeholder": "value",
          "submit": "value",
          "go_back": "value",
          "send_message": "value",
          "text_mode": "value",
          "voice_mode": "value",
          "switched_to_text_mode": "value",
          "switched_to_voice_mode": "value",
          "copy": "value",
          "download": "value",
          "wrap": "value",
          "agent_working": "value",
          "agent_done": "value",
          "agent_error": "value",
          "attach_file": "value",
          "remove_file": "value",
          "file_upload_error": "value",
          "file_type_unsupported": "value",
          "file_too_large": "value",
          "file_limit_reached": "value",
          "typing_indicator": "value",
          "rich_content_unavailable": "value"
        },
        "styles": {
          "base": "value",
          "base_hover": "value",
          "base_active": "value",
          "base_border": "value",
          "base_subtle": "value",
          "base_primary": "value",
          "base_error": "value",
          "accent": "value",
          "accent_hover": "value",
          "accent_active": "value",
          "accent_border": "value",
          "accent_subtle": "value",
          "accent_primary": "value",
          "overlay_padding": 9.99,
          "button_radius": 9.99,
          "input_radius": 9.99,
          "bubble_radius": 9.99,
          "sheet_radius": 9.99,
          "compact_sheet_radius": 9.99,
          "dropdown_sheet_radius": 9.99
        },
        "show_resize_button": true,
        "end_confirmation_enabled": true,
        "mobile_fullscreen": true,
        "transcript_speech_sync_enabled": true,
        "language_selector": true,
        "supports_text_only": true,
        "custom_avatar_path": "value",
        "language_presets": {
          "key": {
            "text_contents": {
              "main_label": "value",
              "start_call": "value",
              "start_chat": "value",
              "new_call": "value",
              "end_call": "value",
              "mute_microphone": "value",
              "change_language": "value",
              "collapse": "value",
              "expand": "value",
              "copied": "value",
              "accept_terms": "value",
              "dismiss_terms": "value",
              "listening_status": "value",
              "speaking_status": "value",
              "connecting_status": "value",
              "chatting_status": "value",
              "queue_waiting_status": "value",
              "queue_waiting_status_short": "value",
              "input_label": "value",
              "input_placeholder": "value",
              "input_placeholder_text_only": "value",
              "input_placeholder_new_conversation": "value",
              "user_ended_conversation": "value",
              "agent_ended_conversation": "value",
              "conversation_id": "value",
              "error_occurred": "value",
              "queue_timed_out": "value",
              "copy_id": "value",
              "initiate_feedback": "value",
              "request_follow_up_feedback": "value",
              "thanks_for_feedback": "value",
              "thanks_for_feedback_details": "value",
              "follow_up_feedback_placeholder": "value",
              "submit": "value",
              "go_back": "value",
              "send_message": "value",
              "text_mode": "value",
              "voice_mode": "value",
              "switched_to_text_mode": "value",
              "switched_to_voice_mode": "value",
              "copy": "value",
              "download": "value",
              "wrap": "value",
              "agent_working": "value",
              "agent_done": "value",
              "agent_error": "value",
              "attach_file": "value",
              "remove_file": "value",
              "file_upload_error": "value",
              "file_type_unsupported": "value",
              "file_too_large": "value",
              "file_limit_reached": "value",
              "typing_indicator": "value",
              "rich_content_unavailable": "value"
            },
            "text_contents_translation": {
              "source": {
                "key": "value"
              },
              "text": {
                "key": "value"
              }
            },
            "terms_text": "value",
            "terms_html": "value",
            "terms_key": "value",
            "terms_translation": {
              "source_hash": "value",
              "text": "value"
            }
          }
        }
      },
      "data_collection": {
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
          "is_omitted": true,
          "name": "value",
          "llm": "gpt-4o-mini"
        }
      },
      "data_collection_scopes": {
        "key": "conversation"
      },
      "analysis_items": {
        "evaluation_criteria": [
          {
            "source": "user",
            "analysis_item_id": "value",
            "version_id": "value",
            "additional_version_ids": [
              "value"
            ],
            "scope": "conversation",
            "weight": 9.99
          }
        ],
        "data_collection": [
          {
            "source": "user",
            "analysis_item_id": "value",
            "version_id": "value",
            "scope": "conversation"
          }
        ]
      },
      "overrides": {
        "conversation_config_override": {
          "asr": {
            "keywords": true
          },
          "turn": {
            "soft_timeout_config": {
              "message": true,
              "additional_soft_timeout_messages": true
            },
            "noisy_environment_message": true
          },
          "tts": {
            "model_id": true,
            "voice_id": true,
            "supported_voices": true,
            "stability": true,
            "speed": true,
            "similarity_boost": true,
            "clone_strength": true,
            "pronunciation_dictionary_locators": true
          },
          "conversation": {
            "text_only": true,
            "max_duration_seconds": true
          },
          "agent": {
            "first_message": true,
            "language": true,
            "max_conversation_duration_message": true,
            "prompt": {
              "prompt": true,
              "llm": true,
              "reasoning_effort": true,
              "tool_ids": true,
              "tool_ids_strict": true,
              "native_mcp_server_ids": true,
              "knowledge_base": true
            }
          }
        },
        "custom_llm_extra_body": true,
        "enable_conversation_initiation_client_data_from_webhook": true,
        "enable_starting_workflow_node_id_from_client": true,
        "enable_procedure_ids_from_client": true
      },
      "workspace_overrides": {
        "conversation_initiation_client_data_webhook": {
          "url": "value",
          "request_headers": {
            "key": "value"
          }
        },
        "webhooks": {
          "post_call_webhook_id": "value",
          "events": [
            "transcript"
          ],
          "transcript_format": "json",
          "exclude_transcript": true,
          "send_audio": true
        }
      },
      "testing": {
        "attached_tests": [
          {
            "test_id": "value",
            "workflow_node_id": "value"
          }
        ]
      },
      "archived": true,
      "guardrails": {
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
      },
      "summary_language": "value",
      "auto_translate_transcript_to_app_language": true,
      "auth": {
        "enable_auth": true,
        "allowlist": [
          {
            "hostname": "value"
          }
        ],
        "require_origin_header": true,
        "shareable_token": "value"
      },
      "call_limits": {
        "agent_concurrency_limit": 123,
        "daily_limit": 123,
        "bursting_enabled": true
      },
      "queueing_config": {
        "enabled": true,
        "wait_timeout_seconds": 123,
        "hold_audio": {
          "audio_path": "value",
          "audio_url": "value",
          "original_filename": "value",
          "duration_secs": 9.99,
          "size_bytes": 123
        }
      },
      "privacy": {
        "record_voice": true,
        "retention_days": 123,
        "delete_transcript_and_pii": true,
        "delete_audio": true,
        "apply_to_existing_conversations": true,
        "zero_retention_mode": true,
        "conversation_history_redaction": {
          "enabled": true,
          "entities": [
            "name"
          ],
          "excluded_data_collection_ids": [
            "value"
          ]
        }
      },
      "trust_context": "unknown",
      "analysis_llm": "gpt-4o-mini",
      "topic_discovery": {},
      "sentiment_analysis": {},
      "alerting": {
        "monitor_configs": {
          "key": {
            "enabled": true,
            "threshold": 0.5,
            "relative_increase_threshold": 9.99,
            "min_failure_count": 123,
            "min_history_bucket_count": 123,
            "min_sample_count": 123,
            "suspect_trigger_threshold": 10,
            "auto_resolve_after_inactive_minutes": 123
          }
        },
        "auto_resolve_after_inactive_minutes": 123,
        "environments": [
          "value"
        ],
        "notifiers": [
          {
            "type": "webhook",
            "webhook_id": "value"
          }
        ]
      }
    },
    "workflow": {
      "edges": {
        "key": {
          "source": "value",
          "target": "value",
          "forward_condition": {
            "label": "value",
            "type": "unconditional"
          },
          "backward_condition": {
            "label": "value",
            "type": "unconditional"
          },
          "reroutes": [
            {
              "x": 9.99,
              "y": 9.99
            }
          ]
        }
      },
      "nodes": {
        "key": {
          "type": "start",
          "position": {
            "x": 9.99,
            "y": 9.99
          },
          "edge_order": [
            "value"
          ]
        }
      },
      "prevent_subagent_loops": true
    }
  },
  "agent_id": "value",
  "branch_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Resubmit(context.Background(), string("value"), &request)
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

type wireRequestAgentsTestsInvocations struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsTestsInvocations records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsTestsInvocations(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsTestsInvocations) {
	t.Helper()
	recorded := make(chan wireRequestAgentsTestsInvocations, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsTestsInvocations{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsTestsInvocations(t *testing.T, recorded <-chan wireRequestAgentsTestsInvocations) wireRequestAgentsTestsInvocations {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsTestsInvocations{}
	}
}

// wireJSONContainsAgentsTestsInvocations reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsTestsInvocations(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsTestsInvocations(got, want)
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

// setWireValueAgentsTestsInvocations sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsTestsInvocations(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsTestsInvocations(got, want []byte) bool {
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
