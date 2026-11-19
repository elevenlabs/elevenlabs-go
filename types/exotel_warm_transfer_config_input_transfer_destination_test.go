package types

import (
	"testing"
)

func TestExotelWarmTransferConfigInputTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(ExotelWarmTransferConfigInputTransferDestination))
}
