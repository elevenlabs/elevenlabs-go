package types

import (
	"testing"
)

func TestQueryParamsJsonSchemaOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(QueryParamsJsonSchemaOutput))
}
