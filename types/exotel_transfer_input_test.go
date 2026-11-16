package types

import (
	"testing"
)

func TestExotelTransferInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "exotel",
  "config": {
    "transfer_destination": {
      "destination": "value",
      "type": "phone"
    },
    "type": "exotel_cold_blind_transfer"
  }
}`
	assertRoundTrip(t, wire, new(ExotelTransferInput))
}
