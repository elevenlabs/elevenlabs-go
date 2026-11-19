package types

import (
	"testing"
)

func TestQueryArrayJsonSchemaPropertyInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "property_kind": "array",
  "description": "value",
  "dynamic_variable": "value",
  "constant_value": [
    {}
  ],
  "is_omitted": true,
  "type": "array",
  "items": {
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
}`
	assertRoundTrip(t, wire, new(QueryArrayJsonSchemaPropertyInput))
}
