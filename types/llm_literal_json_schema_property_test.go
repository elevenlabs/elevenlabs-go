package types

import (
	"testing"
)

func TestLlmLiteralJsonSchemaPropertyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "boolean",
  "description": "value",
  "enum": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(LlmLiteralJsonSchemaProperty))
}
