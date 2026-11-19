package types

import (
	"testing"
)

func TestTwilioWarmConferenceTransferConfigInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(TwilioWarmConferenceTransferConfigInputTransferDestination))
}
