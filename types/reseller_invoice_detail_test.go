package types

import (
	"testing"
)

func TestResellerInvoiceDetailRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "invoice_id": "value",
  "workspace_id": "value",
  "workspace_name": "value",
  "status": "paid",
  "amount_cents": 123,
  "currency": "value",
  "created_unix": 123,
  "paid_unix": 123,
  "pay_url": "value",
  "pdf_url": "value",
  "lines": [
    {
      "kind": "plan",
      "description": "value",
      "quantity": 123,
      "amount_cents": 123,
      "period_start_unix": 123,
      "period_end_unix": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(ResellerInvoiceDetail))
}
