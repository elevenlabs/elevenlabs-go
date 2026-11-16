package types

import (
	"testing"
)

func TestClientToolConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client",
  "name": "value",
  "description": "value",
  "response_timeout_secs": 120,
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
  "expects_response": true,
  "dynamic_variables": {
    "dynamic_variable_placeholders": {
      "key": {}
    }
  },
  "execution_mode": "immediate"
}`
	assertRoundTrip(t, wire, new(ClientToolConfigInput))
}

func TestClientToolConfigInputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client",
  "name": "value",
  "description": "value",
  "response_timeout_secs": 120,
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
  "expects_response": true,
  "dynamic_variables": {
    "dynamic_variable_placeholders": {
      "key": {}
    }
  },
  "execution_mode": "immediate"
}`
	assertRoundTrip(t, wire, new(ClientToolConfigInput))
}
