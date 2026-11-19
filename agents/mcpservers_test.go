package agents

import (
	"testing"
)

func TestMcpServerRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(McpServerRequestModel))
}

func TestMcpServerConfigUpdateRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(McpServerConfigUpdateRequestModel))
}
