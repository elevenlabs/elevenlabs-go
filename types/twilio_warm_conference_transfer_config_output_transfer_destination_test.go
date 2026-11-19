package types

import (
	"testing"
)

func TestTwilioWarmConferenceTransferConfigOutputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(TwilioWarmConferenceTransferConfigOutputTransferDestination))
}
