package types

import (
	"testing"
)

func TestSipColdConferenceTransferConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "type": "sip_cold_conference_transfer"
}`
	assertRoundTrip(t, wire, new(SipColdConferenceTransferConfigInput))
}
