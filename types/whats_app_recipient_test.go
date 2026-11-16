package types

import (
	"testing"
)

func TestWhatsAppRecipientRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone": "value",
  "bsuid": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppRecipient))
}
