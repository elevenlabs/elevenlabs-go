package types

import (
	"testing"
)

func TestContentSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "string",
  "enum": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ContentSchema))
}
