package productions

import (
	"testing"
)

func TestCreateOrderRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sandbox": true
}`
	assertRoundTrip(t, wire, new(CreateOrderRequest))
}

func TestBodyUpdateOrderV1ProductionsOrdersOrderIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "request": {
    "name": "value"
  }
}`
	assertRoundTrip(t, wire, new(BodyUpdateOrderV1ProductionsOrdersOrderIdPatch))
}
