package types

import (
	"testing"
)

func TestTwilioColdBlindTransferConfigOutputRoundTrip(t *testing.T) {
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
  "type": "twilio_cold_blind_transfer"
}`
	assertRoundTrip(t, wire, new(TwilioColdBlindTransferConfigOutput))
}
