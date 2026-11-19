package types

import (
	"testing"
)

func TestUserDataCollectionAnalysisItemOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "data_collection",
  "value_schema": {
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
	assertRoundTrip(t, wire, new(UserDataCollectionAnalysisItemOutput))
}
