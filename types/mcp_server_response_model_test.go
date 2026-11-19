package types

import (
	"testing"
)

func TestMcpServerResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(McpServerResponseModel))
}

func TestMcpServerResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(McpServerResponseModel))
}

func TestMcpServerResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(McpServerResponseModel))
}
