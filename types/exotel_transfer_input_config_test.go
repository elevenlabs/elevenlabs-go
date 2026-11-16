package types

import (
	"testing"
)

func TestExotelTransferInputConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_destination": {
    "destination": "value",
    "type": "phone"
  },
  "type": "exotel_cold_blind_transfer"
}`
	assertRoundTrip(t, wire, new(ExotelTransferInputConfig))
}
