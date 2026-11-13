package types

import (
	"testing"
)

func TestLlmArrayJsonSchemaPropertyOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "array",
  "description": "value",
  "items": {
    "type": "boolean",
    "description": "value",
    "enum": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(LlmArrayJsonSchemaPropertyOutput))
}
