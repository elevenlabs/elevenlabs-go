package types

import (
	"testing"
)

func TestAstllmNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "value_schema": {
    "type": "boolean",
    "description": "value",
    "enum": [
      "value"
    ]
  },
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(AstllmNodeOutput))
}
