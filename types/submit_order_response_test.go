package types

import (
	"testing"
)

func TestSubmitOrderResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "state": "open",
  "submitted_at": "2024-01-01T00:00:00Z"
}`
	assertRoundTrip(t, wire, new(SubmitOrderResponse))
}
