package toolconfigs

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

	mcpservers "github.com/elevenlabs/elevenlabs-go/agents/mcpservers"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversToolconfigs(t, 200, `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request mcpservers.McpToolConfigOverrideCreateRequestModel
	if err := json.Unmarshal([]byte(`{
  "environment": "value",
  "mcp_server_id": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
  ],
  "tool_name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	environmentValue := string("value")
	request.Environment = &environmentValue
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversToolconfigs(gotResponse, []byte(`{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)
	}

	req := receiveWireAgentsMcpserversToolconfigs(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/mcp-servers/value/tool-configs" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value/tool-configs")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMcpserversToolconfigs(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError409(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversToolconfigs(t, 409, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request mcpservers.McpToolConfigOverrideCreateRequestModel
	if err := json.Unmarshal([]byte(`{
  "environment": "value",
  "mcp_server_id": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
  ],
  "tool_name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	environmentValue := string("value")
	request.Environment = &environmentValue
	_, err := client.Create(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 409 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 409 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 409)
	}
}

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversToolconfigs(t, 200, `{
  "tool_name": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversToolconfigs(gotResponse, []byte(`{
  "tool_name": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tool_name": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
}`)
	}

	req := receiveWireAgentsMcpserversToolconfigs(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/mcp-servers/value/tool-configs/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value/tool-configs/value2")
	}
}

func TestGetWireError404(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversToolconfigs(t, 404, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Get(context.Background(), string("value"), string("value2"))
	if err == nil {
		t.Fatal("expected an error for a 404 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 404 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 404)
	}
}

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversToolconfigs(t, 200, `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request mcpservers.McpToolConfigOverrideUpdateRequestModel
	if err := json.Unmarshal([]byte(`{
  "environment": "value",
  "tool_name": "value",
  "mcp_server_id": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	environmentValue := string("value")
	request.Environment = &environmentValue
	result, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversToolconfigs(gotResponse, []byte(`{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)
	}

	req := receiveWireAgentsMcpserversToolconfigs(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/mcp-servers/value/tool-configs/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value/tool-configs/value2")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMcpserversToolconfigs(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError404(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversToolconfigs(t, 404, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request mcpservers.McpToolConfigOverrideUpdateRequestModel
	if err := json.Unmarshal([]byte(`{
  "environment": "value",
  "tool_name": "value",
  "mcp_server_id": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
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
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	environmentValue := string("value")
	request.Environment = &environmentValue
	_, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err == nil {
		t.Fatal("expected an error for a 404 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 404 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 404)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversToolconfigs(t, 200, `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversToolconfigs(gotResponse, []byte(`{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "config": {
    "approval_policy": "auto_approve_all",
    "tool_approval_hashes": [
      {
        "tool_name": "value",
        "tool_hash": "value",
        "approval_policy": "auto_approved"
      }
    ],
    "transport": "SSE",
    "url": "value",
    "secret_token": {
      "secret_id": "value"
    },
    "request_headers": {
      "key": "value"
    },
    "request_meta": {
      "key": "value"
    },
    "auth_connection": {
      "auth_connection_id": "value"
    },
    "name": "value",
    "description": "value",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "execution_mode": "immediate",
    "response_timeout_secs": 123,
    "tool_config_overrides": [
      {
        "tool_name": "value",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "execution_mode": "immediate",
        "response_timeout_secs": 123,
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "input_overrides": {
          "key": {
            "source": "constant",
            "constant_value": "value"
          }
        },
        "response_mocks": [
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
    ],
    "disable_compression": true
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "dependent_agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "metadata": {
    "created_at": 123,
    "owner_user_id": "value"
  }
}`)
	}

	req := receiveWireAgentsMcpserversToolconfigs(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/mcp-servers/value/tool-configs/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value/tool-configs/value2")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversToolconfigs(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"), string("value2"))
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

type wireRequestAgentsMcpserversToolconfigs struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsMcpserversToolconfigs records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsMcpserversToolconfigs(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsMcpserversToolconfigs) {
	t.Helper()
	recorded := make(chan wireRequestAgentsMcpserversToolconfigs, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsMcpserversToolconfigs{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsMcpserversToolconfigs(t *testing.T, recorded <-chan wireRequestAgentsMcpserversToolconfigs) wireRequestAgentsMcpserversToolconfigs {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsMcpserversToolconfigs{}
	}
}

// wireJSONContainsAgentsMcpserversToolconfigs reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsMcpserversToolconfigs(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsMcpserversToolconfigs(got, want)
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

// setWireValueAgentsMcpserversToolconfigs sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsMcpserversToolconfigs(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsMcpserversToolconfigs(got, want []byte) bool {
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
