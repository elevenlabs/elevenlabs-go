package types

import (
	"testing"
)

func TestDataCollectionResultCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "data_collection_id": "value",
  "name": "value",
  "value": {},
  "json_schema": {
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
  },
  "rationale": "value"
}`
	assertRoundTrip(t, wire, new(DataCollectionResultCommonModel))
}
