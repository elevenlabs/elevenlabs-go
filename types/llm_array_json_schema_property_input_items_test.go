package types

import (
	"testing"
)

func TestLlmArrayJsonSchemaPropertyInputItemsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "boolean",
  "description": "value",
  "enum": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(LlmArrayJsonSchemaPropertyInputItems))
}
