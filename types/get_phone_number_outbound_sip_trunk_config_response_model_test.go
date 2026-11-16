package types

import (
	"testing"
)

func TestGetPhoneNumberOutboundSipTrunkConfigResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "address": "value",
  "transport": "auto",
  "media_encryption": "disabled",
  "headers": {
    "key": "value"
  },
  "attributes_to_headers": {
    "key": "value"
  },
  "has_auth_credentials": true,
  "username": "value",
  "has_outbound_trunk": true,
  "enabled_codecs": [
    "G722/8000"
  ]
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberOutboundSipTrunkConfigResponseModel))
}
