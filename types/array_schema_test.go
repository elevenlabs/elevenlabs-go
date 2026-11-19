package types

import (
	"testing"
)

func TestArraySchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "array",
  "items": {
    "title": "value",
    "description": "value",
    "type": "string",
    "enum": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(ArraySchema))
}
