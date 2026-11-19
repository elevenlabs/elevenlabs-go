package types

import (
	"testing"
)

func TestPhoneDestinationOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone"
}`
	assertRoundTrip(t, wire, new(PhoneDestinationOutput))
}
