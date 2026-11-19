package types

import (
	"testing"
)

func TestObjectSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "object",
  "properties": {
    "key": {
      "title": "value",
      "description": "value",
      "type": "string",
      "enum": [
        "value"
      ]
    }
  },
  "required": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ObjectSchema))
}
