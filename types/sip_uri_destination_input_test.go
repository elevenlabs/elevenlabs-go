package types

import (
	"testing"
)

func TestSipUriDestinationInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "sip_uri"
}`
	assertRoundTrip(t, wire, new(SipUriDestinationInput))
}
