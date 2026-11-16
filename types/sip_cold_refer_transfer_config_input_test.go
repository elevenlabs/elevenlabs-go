package types

import (
	"testing"
)

func TestSipColdReferTransferConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "custom_sip_headers": [
    {
      "type": "static",
      "key": "value",
      "value": "value"
    }
  ],
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "type": "sip_cold_refer_transfer",
  "uui": {
    "data": "value",
    "protocol_discriminator": "00",
    "protocol_discriminator_mode": "prefix"
  },
  "play_dialtone": true
}`
	assertRoundTrip(t, wire, new(SipColdReferTransferConfigInput))
}
