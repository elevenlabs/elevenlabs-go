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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversClient(t, 200, `{
  "mcp_servers": [
    {
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
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversClient(gotResponse, []byte(`{
  "mcp_servers": [
    {
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
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "mcp_servers": [
    {
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
    }
  ]
}`)
	}

	req := receiveWireAgentsMcpserversClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/mcp-servers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversClient(t, 422, `{
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

	_, err := client.List(context.Background())
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversClient(t, 200, `{
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

	var request agents.McpServerRequestModel
	if err := json.Unmarshal([]byte(`{
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
  }
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
	if !wireJSONBytesEqualAgentsMcpserversClient(gotResponse, []byte(`{
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

	req := receiveWireAgentsMcpserversClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/mcp-servers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMcpserversClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversClient(t, 422, `{
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

	var request agents.McpServerRequestModel
	if err := json.Unmarshal([]byte(`{
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
  }
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversClient(t, 200, `{
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

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsMcpserversClient(gotResponse, []byte(`{
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

	req := receiveWireAgentsMcpserversClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/mcp-servers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversClient(t, 422, `{
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

	server, recorded := serveWireAgentsMcpserversClient(t, 200, `{
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

	var request agents.McpServerConfigUpdateRequestModel
	if err := json.Unmarshal([]byte(`{
  "mcp_server_id": "value",
  "approval_policy": "auto_approve_all",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "request_headers": {
    "key": "value"
  },
  "request_meta": {
    "key": "value"
  },
  "disable_compression": true,
  "secret_token": {
    "secret_id": "value"
  },
  "auth_connection": {
    "auth_connection_id": "value"
  }
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
	if !wireJSONBytesEqualAgentsMcpserversClient(gotResponse, []byte(`{
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

	req := receiveWireAgentsMcpserversClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/mcp-servers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsMcpserversClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversClient(t, 422, `{
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

	var request agents.McpServerConfigUpdateRequestModel
	if err := json.Unmarshal([]byte(`{
  "mcp_server_id": "value",
  "approval_policy": "auto_approve_all",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "request_headers": {
    "key": "value"
  },
  "request_meta": {
    "key": "value"
  },
  "disable_compression": true,
  "secret_token": {
    "secret_id": "value"
  },
  "auth_connection": {
    "auth_connection_id": "value"
  }
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsMcpserversClient(t, 200, `{}`)

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
	if !wireJSONBytesEqualAgentsMcpserversClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsMcpserversClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/mcp-servers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/mcp-servers/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsMcpserversClient(t, 422, `{
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

type wireRequestAgentsMcpserversClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsMcpserversClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsMcpserversClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsMcpserversClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsMcpserversClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsMcpserversClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsMcpserversClient(t *testing.T, recorded <-chan wireRequestAgentsMcpserversClient) wireRequestAgentsMcpserversClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsMcpserversClient{}
	}
}

// wireJSONContainsAgentsMcpserversClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsMcpserversClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsMcpserversClient(got, want)
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

// setWireValueAgentsMcpserversClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsMcpserversClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsMcpserversClient(got, want []byte) bool {
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
