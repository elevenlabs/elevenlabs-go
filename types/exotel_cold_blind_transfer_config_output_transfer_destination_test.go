package types

import (
	"testing"
)

func TestExotelColdBlindTransferConfigOutputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(ExotelColdBlindTransferConfigOutputTransferDestination))
}
