package types

import (
	"testing"
)

func TestSipColdReferTransferConfigInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(SipColdReferTransferConfigInputTransferDestination))
}
