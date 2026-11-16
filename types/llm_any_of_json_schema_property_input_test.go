package types

import (
	"testing"
)

func TestLlmAnyOfJsonSchemaPropertyInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "any_of": [
    {
      "type": "object",
      "required": [
        "value"
      ],
      "description": "value",
      "properties": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ]
        }
      }
    }
  ],
  "description": "value"
}`
	assertRoundTrip(t, wire, new(LlmAnyOfJsonSchemaPropertyInput))
}
