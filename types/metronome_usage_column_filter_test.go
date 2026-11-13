package types

import (
	"testing"
)

func TestMetronomeUsageColumnFilterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "column": "model",
  "operation": "eq",
  "values": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(MetronomeUsageColumnFilter))
}
