package types

import (
	"testing"
)

func TestTransferToNumberToolConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "transfer_to_number",
  "transfers": [
    {
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
      "phone_number": "value",
      "condition": "value"
    }
  ],
  "transfer_rules": [
    {
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
    }
  ],
  "enable_client_message": true
}`
	assertRoundTrip(t, wire, new(TransferToNumberToolConfigOutput))
}
