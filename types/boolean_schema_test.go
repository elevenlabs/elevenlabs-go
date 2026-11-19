package types

import (
	"testing"
)

func TestBooleanSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "boolean"
}`
	assertRoundTrip(t, wire, new(BooleanSchema))
}
