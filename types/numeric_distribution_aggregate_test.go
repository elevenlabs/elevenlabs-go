package types

import (
	"testing"
)

func TestNumericDistributionAggregateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "count": 123,
  "sum": 9.99,
  "min": 9.99,
  "max": 9.99
}`
	assertRoundTrip(t, wire, new(NumericDistributionAggregate))
}
