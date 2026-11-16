package types

import (
	"testing"
)

func TestResellerInvoicesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "invoices": [
    {
      "invoice_id": "value",
      "workspace_id": "value",
      "workspace_name": "value",
      "status": "paid",
      "amount_cents": 123,
      "currency": "value",
      "created_unix": 123,
      "paid_unix": 123,
      "pay_url": "value",
      "pdf_url": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(ResellerInvoicesResponse))
}
