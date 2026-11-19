package types

import (
	"testing"
)

func TestWhatsAppOutboundCallResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "success": true,
  "message": "value",
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppOutboundCallResponse))
}
