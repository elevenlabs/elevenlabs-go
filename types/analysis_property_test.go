package types

import (
	"testing"
)

func TestAnalysisPropertyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  "is_omitted": true,
  "name": "value",
  "llm": "gpt-4o-mini"
}`
	assertRoundTrip(t, wire, new(AnalysisProperty))
}
