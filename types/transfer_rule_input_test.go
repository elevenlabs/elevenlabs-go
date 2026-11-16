package types

import (
	"testing"
)

func TestTransferRuleInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "provider_configs": [
    {
      "type": "twilio",
      "config": {
        "transfer_destination": {
          "destination": "value",
          "type": "phone"
        },
        "post_dial_digits": {
          "type": "static",
          "value": "0"
        },
        "type": "twilio_warm_conference_transfer"
      }
    }
  ],
  "condition": "value"
}`
	assertRoundTrip(t, wire, new(TransferRuleInput))
}
