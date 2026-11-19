package types

import (
	"testing"
)

func TestTwilioColdBlindTransferConfigInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(TwilioColdBlindTransferConfigInputTransferDestination))
}
