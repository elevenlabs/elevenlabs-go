package types

import (
	"testing"
)

func TestInvoiceResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "amount_due_cents": 123,
  "subtotal_cents": 123,
  "tax_cents": 123,
  "discount_percent_off": 9.99,
  "discount_amount_off": 9.99,
  "discounts": [
    {
      "discount_percent_off": 9.99,
      "discount_amount_off": 9.99
    }
  ],
  "next_payment_attempt_unix": 123,
  "payment_intent_status": "canceled",
  "payment_intent_statusses": [
    "canceled"
  ]
}`
	assertRoundTrip(t, wire, new(InvoiceResponse))
}
