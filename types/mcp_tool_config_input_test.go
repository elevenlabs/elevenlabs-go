package types

import (
	"testing"
)

func TestMcpToolConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
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
  "integration_type": "mcp_server",
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
  },
  "approval_policy": "auto_approve_all",
  "mcp_tool_name": "value",
  "mcp_tool_description": "value",
  "mcp_server_id": "value",
  "mcp_server_name": "value",
  "mcp_input_schema": {
    "key": {}
  },
  "execution_mode": "immediate",
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(McpToolConfigInput))
}

func TestMcpToolConfigInputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
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
  "integration_type": "mcp_server",
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
  },
  "approval_policy": "auto_approve_all",
  "mcp_tool_name": "value",
  "mcp_tool_description": "value",
  "mcp_server_id": "value",
  "mcp_server_name": "value",
  "mcp_input_schema": {
    "key": {}
  },
  "execution_mode": "immediate",
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(McpToolConfigInput))
}
