package types

import (
	"testing"
)

func TestResellerInvoiceSummaryRoundTrip(t *testing.T) {
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
  "pdf_url": "value"
}`
	assertRoundTrip(t, wire, new(ResellerInvoiceSummary))
}
