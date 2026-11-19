package types

import (
	"testing"
)

func TestPhoneDestinationInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(PhoneDestinationInput))
}
