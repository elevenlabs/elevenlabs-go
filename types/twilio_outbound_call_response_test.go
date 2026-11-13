package types

import (
	"testing"
)

func TestTwilioOutboundCallResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "success": true,
  "message": "value",
  "conversation_id": "value",
  "callSid": "value"
}`
	assertRoundTrip(t, wire, new(TwilioOutboundCallResponse))
}
