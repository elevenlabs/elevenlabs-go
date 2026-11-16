package types

import (
	"testing"
)

func TestIntegerSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "integer"
}`
	assertRoundTrip(t, wire, new(IntegerSchema))
}
