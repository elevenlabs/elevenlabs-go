package types

import (
	"testing"
)

func TestDiscountResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "discount_percent_off": 9.99,
  "discount_amount_off": 9.99
}`
	assertRoundTrip(t, wire, new(DiscountResponseModel))
}
