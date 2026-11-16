package types

import (
	"testing"
)

func TestTwilioTransferOutputConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "post_dial_digits": {
    "type": "static",
    "value": "0"
  },
  "type": "twilio_warm_conference_transfer"
}`
	assertRoundTrip(t, wire, new(TwilioTransferOutputConfig))
}
