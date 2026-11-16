package types

import (
	"testing"
)

func TestUpdateOrderResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(UpdateOrderResponse))
}
