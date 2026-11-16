package types

import (
	"testing"
)

func TestInboundSipTrunkConfigRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(InboundSipTrunkConfigRequestModel))
}
