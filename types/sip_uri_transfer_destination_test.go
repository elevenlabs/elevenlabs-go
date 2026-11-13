package types

import (
	"testing"
)

func TestSipUriTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "sip_uri",
  "sip_uri": "value"
}`
	assertRoundTrip(t, wire, new(SipUriTransferDestination))
}
