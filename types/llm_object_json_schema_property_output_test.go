package types

import (
	"testing"
)

func TestLlmObjectJsonSchemaPropertyOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(LlmObjectJsonSchemaPropertyOutput))
}
