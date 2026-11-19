package types

import (
	"testing"
)

func TestPhoneNumberTransferTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "phone",
  "phone_number": "value"
}`
	assertRoundTrip(t, wire, new(PhoneNumberTransferTransferDestination))
}
