package types

import (
	"testing"
)

func TestSipColdConferenceTransferConfigOutputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(SipColdConferenceTransferConfigOutputTransferDestination))
}
