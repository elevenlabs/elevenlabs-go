package types

import (
	"testing"
)

func TestWhatsAppOutboundMessageResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppOutboundMessageResponse))
}
