package batchcalls

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

	server, recorded := serveWireAgentsBatchcalls(t, 200, `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost
	if err := json.Unmarshal([]byte(`{
  "call_name": "value",
  "agent_id": "value",
  "recipients": [
    {
      "id": "value",
      "phone_number": "value",
      "whatsapp_user_id": "value",
      "conversation_initiation_client_data": {
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
        }
      }
    }
  ],
  "scheduled_time_unix": 123,
  "agent_phone_number_id": "value",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "timezone": "value",
  "branch_id": "value",
  "environment": "value",
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsBatchcalls(gotResponse, []byte(`{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/batch-calling/submit" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/submit")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsBatchcalls(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

	var request agents.BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost
	if err := json.Unmarshal([]byte(`{
  "call_name": "value",
  "agent_id": "value",
  "recipients": [
    {
      "id": "value",
      "phone_number": "value",
      "whatsapp_user_id": "value",
      "conversation_initiation_client_data": {
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
        }
      }
    }
  ],
  "scheduled_time_unix": 123,
  "agent_phone_number_id": "value",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "timezone": "value",
  "branch_id": "value",
  "environment": "value",
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsBatchcalls(t, 200, `{
  "batch_calls": [
    {
      "id": "value",
      "phone_number_id": "value",
      "phone_provider": "twilio",
      "whatsapp_params": {
        "whatsapp_phone_number_id": "value",
        "whatsapp_call_permission_request_template_name": "value",
        "whatsapp_call_permission_request_template_language_code": "value"
      },
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "environment": "value",
      "created_at_unix": 123,
      "scheduled_time_unix": 123,
      "timezone": "value",
      "total_calls_dispatched": 123,
      "total_calls_scheduled": 123,
      "total_calls_finished": 123,
      "last_updated_at_unix": 123,
      "status": "pending",
      "retry_count": 123,
      "telephony_call_config": {
        "ringing_timeout_secs": 123,
        "twilio_call_recording_enabled": true,
        "twilio_machine_detection": {
          "mode": "enable",
          "actions": [
            {
              "action": "hang_up",
              "answered_by": [
                "human"
              ]
            }
          ]
        }
      },
      "target_concurrency_limit": 123,
      "agent_name": "value",
      "branch_name": "value"
    }
  ],
  "next_doc": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	limitValue := int(123)
	lastDocValue := string("value")
	agentIdValue := string("value")
	request := &agents.ListBatchCallsRequest{
		Limit:   &limitValue,
		LastDoc: &lastDocValue,
		AgentId: &agentIdValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsBatchcalls(gotResponse, []byte(`{
  "batch_calls": [
    {
      "id": "value",
      "phone_number_id": "value",
      "phone_provider": "twilio",
      "whatsapp_params": {
        "whatsapp_phone_number_id": "value",
        "whatsapp_call_permission_request_template_name": "value",
        "whatsapp_call_permission_request_template_language_code": "value"
      },
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "environment": "value",
      "created_at_unix": 123,
      "scheduled_time_unix": 123,
      "timezone": "value",
      "total_calls_dispatched": 123,
      "total_calls_scheduled": 123,
      "total_calls_finished": 123,
      "last_updated_at_unix": 123,
      "status": "pending",
      "retry_count": 123,
      "telephony_call_config": {
        "ringing_timeout_secs": 123,
        "twilio_call_recording_enabled": true,
        "twilio_machine_detection": {
          "mode": "enable",
          "actions": [
            {
              "action": "hang_up",
              "answered_by": [
                "human"
              ]
            }
          ]
        }
      },
      "target_concurrency_limit": 123,
      "agent_name": "value",
      "branch_name": "value"
    }
  ],
  "next_doc": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "batch_calls": [
    {
      "id": "value",
      "phone_number_id": "value",
      "phone_provider": "twilio",
      "whatsapp_params": {
        "whatsapp_phone_number_id": "value",
        "whatsapp_call_permission_request_template_name": "value",
        "whatsapp_call_permission_request_template_language_code": "value"
      },
      "name": "value",
      "agent_id": "value",
      "branch_id": "value",
      "environment": "value",
      "created_at_unix": 123,
      "scheduled_time_unix": 123,
      "timezone": "value",
      "total_calls_dispatched": 123,
      "total_calls_scheduled": 123,
      "total_calls_finished": 123,
      "last_updated_at_unix": 123,
      "status": "pending",
      "retry_count": 123,
      "telephony_call_config": {
        "ringing_timeout_secs": 123,
        "twilio_call_recording_enabled": true,
        "twilio_machine_detection": {
          "mode": "enable",
          "actions": [
            {
              "action": "hang_up",
              "answered_by": [
                "human"
              ]
            }
          ]
        }
      },
      "target_concurrency_limit": 123,
      "agent_name": "value",
      "branch_name": "value"
    }
  ],
  "next_doc": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/batch-calling/workspace" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/workspace")
	}
	if got := req.Query.Get("limit"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "limit", got, "123")
	}
	if got := req.Query.Get("last_doc"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "last_doc", got, "value")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

	limitValue := int(123)
	lastDocValue := string("value")
	agentIdValue := string("value")
	request := &agents.ListBatchCallsRequest{
		Limit:   &limitValue,
		LastDoc: &lastDocValue,
		AgentId: &agentIdValue,
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

	server, recorded := serveWireAgentsBatchcalls(t, 200, `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value",
  "recipients": [
    {
      "id": "value",
      "phone_number": "value",
      "whatsapp_user_id": "value",
      "status": "pending",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "conversation_id": "value",
      "conversation_initiation_client_data": {
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
	if !wireJSONBytesEqualAgentsBatchcalls(gotResponse, []byte(`{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value",
  "recipients": [
    {
      "id": "value",
      "phone_number": "value",
      "whatsapp_user_id": "value",
      "status": "pending",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "conversation_id": "value",
      "conversation_initiation_client_data": {
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
      }
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value",
  "recipients": [
    {
      "id": "value",
      "phone_number": "value",
      "whatsapp_user_id": "value",
      "status": "pending",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "conversation_id": "value",
      "conversation_initiation_client_data": {
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
      }
    }
  ]
}`)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/batch-calling/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsBatchcalls(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/batch-calling/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

	err := client.Delete(context.Background(), string("value"))
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

	server, recorded := serveWireAgentsBatchcalls(t, 200, `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Cancel(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsBatchcalls(gotResponse, []byte(`{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/batch-calling/value/cancel" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/value/cancel")
	}
}

func TestCancelWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

	_, err := client.Cancel(context.Background(), string("value"))
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

func TestRetryWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsBatchcalls(t, 200, `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Retry(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsBatchcalls(gotResponse, []byte(`{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "phone_number_id": "value",
  "phone_provider": "twilio",
  "whatsapp_params": {
    "whatsapp_phone_number_id": "value",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  },
  "name": "value",
  "agent_id": "value",
  "branch_id": "value",
  "environment": "value",
  "created_at_unix": 123,
  "scheduled_time_unix": 123,
  "timezone": "value",
  "total_calls_dispatched": 123,
  "total_calls_scheduled": 123,
  "total_calls_finished": 123,
  "last_updated_at_unix": 123,
  "status": "pending",
  "retry_count": 123,
  "telephony_call_config": {
    "ringing_timeout_secs": 123,
    "twilio_call_recording_enabled": true,
    "twilio_machine_detection": {
      "mode": "enable",
      "actions": [
        {
          "action": "hang_up",
          "answered_by": [
            "human"
          ]
        }
      ]
    }
  },
  "target_concurrency_limit": 123,
  "agent_name": "value",
  "branch_name": "value"
}`)
	}

	req := receiveWireAgentsBatchcalls(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/batch-calling/value/retry" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/batch-calling/value/retry")
	}
}

func TestRetryWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsBatchcalls(t, 422, `{
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

	_, err := client.Retry(context.Background(), string("value"))
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

type wireRequestAgentsBatchcalls struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsBatchcalls records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsBatchcalls(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsBatchcalls) {
	t.Helper()
	recorded := make(chan wireRequestAgentsBatchcalls, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsBatchcalls{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsBatchcalls(t *testing.T, recorded <-chan wireRequestAgentsBatchcalls) wireRequestAgentsBatchcalls {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsBatchcalls{}
	}
}

// wireJSONContainsAgentsBatchcalls reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsBatchcalls(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsBatchcalls(got, want)
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

// setWireValueAgentsBatchcalls sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsBatchcalls(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsBatchcalls(got, want []byte) bool {
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
