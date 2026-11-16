package types

import (
	"testing"
)

func TestUpdateOrderRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(UpdateOrderRequest))
}
