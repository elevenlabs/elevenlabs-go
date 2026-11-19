package types

import (
	"testing"
)

func TestSubAgentInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(SubAgentInput))
}
