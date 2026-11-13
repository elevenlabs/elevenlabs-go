package types

import (
	"testing"
)

func TestOutboundSipTrunkConfigRequestModelRoundTrip(t *testing.T) {
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
  "credentials": {
    "username": "value",
    "password": "value"
  },
  "enabled_codecs": [
    "G722/8000"
  ]
}`
	assertRoundTrip(t, wire, new(OutboundSipTrunkConfigRequestModel))
}
