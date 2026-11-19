package types

import (
	"testing"
)

func TestAstllmNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "value_schema": {
    "type": "boolean",
    "description": "value",
    "enum": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(AstllmNodeInput))
}
