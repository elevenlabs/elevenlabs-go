package types

import (
	"testing"
)

func TestModelRatesResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "character_cost_multiplier": 9.99,
  "cost_discount_multiplier": 9.99
}`
	assertRoundTrip(t, wire, new(ModelRatesResponseModel))
}
