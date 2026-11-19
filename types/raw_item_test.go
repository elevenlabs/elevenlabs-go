package types

import (
	"testing"
)

func TestRawItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "description": "value",
  "price": 9.99,
  "duration_minutes": 123
}`
	assertRoundTrip(t, wire, new(RawItem))
}
