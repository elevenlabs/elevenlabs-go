package types

import (
	"testing"
)

func TestOrderSummaryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "name": "value",
  "state": "open",
  "total_amount_usd": 9.99,
  "sandbox": true,
  "submitted_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "cancel_reason": "value"
}`
	assertRoundTrip(t, wire, new(OrderSummary))
}
