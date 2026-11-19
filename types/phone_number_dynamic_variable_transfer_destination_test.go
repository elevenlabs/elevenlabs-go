package types

import (
	"testing"
)

func TestPhoneNumberDynamicVariableTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "phone_dynamic_variable",
  "phone_number": "value"
}`
	assertRoundTrip(t, wire, new(PhoneNumberDynamicVariableTransferDestination))
}
