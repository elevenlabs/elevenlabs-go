package types

import (
	"testing"
)

func TestResellerInvoiceLineRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "plan",
  "description": "value",
  "quantity": 123,
  "amount_cents": 123,
  "period_start_unix": 123,
  "period_end_unix": 123
}`
	assertRoundTrip(t, wire, new(ResellerInvoiceLine))
}
