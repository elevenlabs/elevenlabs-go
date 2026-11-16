package types

import (
	"testing"
)

func TestWhatsAppOpeningMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "template_name": "value",
  "wamid": "value",
  "sent_at_unix": 123,
  "delivered_at_unix": 123,
  "read_at_unix": 123,
  "replied_at_unix": 123,
  "failed_at_unix": 123,
  "error_code": "value",
  "pricing_category": "value",
  "billable": true
}`
	assertRoundTrip(t, wire, new(WhatsAppOpeningMessage))
}
