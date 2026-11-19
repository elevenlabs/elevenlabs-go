package types

import (
	"testing"
)

func TestCreateSipTrunkPhoneNumberRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "sip_trunk",
  "agent_id": "value",
  "inbound_trunk_config": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "outbound_trunk_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "enabled_codecs": [
      "G722/8000"
    ]
  }
}`
	assertRoundTrip(t, wire, new(CreateSipTrunkPhoneNumberRequest))
}
