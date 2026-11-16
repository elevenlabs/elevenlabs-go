package types

import (
	"testing"
)

func TestExotelWarmTransferConfigOutputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(ExotelWarmTransferConfigOutputTransferDestination))
}
