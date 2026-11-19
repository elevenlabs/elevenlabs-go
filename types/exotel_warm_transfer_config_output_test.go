package types

import (
	"testing"
)

func TestExotelWarmTransferConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "type": "exotel_warm_transfer"
}`
	assertRoundTrip(t, wire, new(ExotelWarmTransferConfigOutput))
}
