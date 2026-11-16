package types

import (
	"testing"
)

func TestSipTransferOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "sip_trunk",
  "config": {
    "transfer_destination": {
      "destination": "value",
      "type": "phone"
    },
    "type": "sip_cold_conference_transfer"
  }
}`
	assertRoundTrip(t, wire, new(SipTransferOutput))
}
