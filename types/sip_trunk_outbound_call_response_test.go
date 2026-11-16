package types

import (
	"testing"
)

func TestSipTrunkOutboundCallResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "success": true,
  "message": "value",
  "conversation_id": "value",
  "sip_call_id": "value"
}`
	assertRoundTrip(t, wire, new(SipTrunkOutboundCallResponse))
}
