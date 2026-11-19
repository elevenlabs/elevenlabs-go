package types

import (
	"testing"
)

func TestPromptAgentApiModelOutputToolsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelOutputToolsItem))
}
