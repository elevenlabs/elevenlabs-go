package types

import (
	"testing"
)

func TestSipUriDestinationOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "sip_uri"
}`
	assertRoundTrip(t, wire, new(SipUriDestinationOutput))
}
