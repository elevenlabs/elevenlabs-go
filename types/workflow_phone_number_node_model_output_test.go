package types

import (
	"testing"
)

func TestWorkflowPhoneNumberNodeModelOutputRoundTrip(t *testing.T) {
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
    "type": "phone",
    "phone_number": "value"
  },
  "transfer_type": "blind",
  "sip_refer_play_dialtone": true,
  "uui": {
    "data": "value",
    "protocol_discriminator": "00",
    "protocol_discriminator_mode": "prefix"
  },
  "post_dial_digits": {
    "type": "static",
    "value": "0"
  },
  "type": "phone_number",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(WorkflowPhoneNumberNodeModelOutput))
}
