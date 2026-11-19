package types

import (
	"testing"
)

func TestCreateOrderResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "sandbox": true
}`
	assertRoundTrip(t, wire, new(CreateOrderResponse))
}
