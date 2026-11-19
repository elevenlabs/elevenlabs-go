package types

import (
	"testing"
)

func TestSipColdReferTransferConfigOutputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(SipColdReferTransferConfigOutputTransferDestination))
}
