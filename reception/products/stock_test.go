package products

import (
	"testing"
)

func TestReceptionUpdateProductStockRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "location_id": "value",
  "on_hand": 123,
  "expected_on_hand": 123
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateProductStockRequest))
}

func TestReceptionAdjustProductStockRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "location_id": "value",
  "delta": 123
}`
	assertRoundTrip(t, wire, new(ReceptionAdjustProductStockRequest))
}
