package types

import (
	"testing"
)

func TestNumberSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "number",
  "enum": [
    9.99
  ]
}`
	assertRoundTrip(t, wire, new(NumberSchema))
}
