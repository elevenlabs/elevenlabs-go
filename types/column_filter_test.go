package types

import (
	"testing"
)

func TestColumnFilterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "column": "value",
  "operation": "in",
  "values": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ColumnFilter))
}

func TestColumnFilterRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "column": "value",
  "operation": "in",
  "values": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ColumnFilter))
}
