package types

import (
	"testing"
)

func TestSipColdConferenceTransferConfigInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(SipColdConferenceTransferConfigInputTransferDestination))
}
