package types

import (
	"testing"
)

func TestExotelWarmTransferConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "type": "exotel_warm_transfer"
}`
	assertRoundTrip(t, wire, new(ExotelWarmTransferConfigInput))
}
