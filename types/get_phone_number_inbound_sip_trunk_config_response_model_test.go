package types

import (
	"testing"
)

func TestGetPhoneNumberInboundSipTrunkConfigResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "allowed_addresses": [
    "value"
  ],
  "allowed_numbers": [
    "value"
  ],
  "media_encryption": "disabled",
  "has_auth_credentials": true,
  "username": "value",
  "remote_domains": [
    "value"
  ],
  "attributes_to_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberInboundSipTrunkConfigResponseModel))
}
