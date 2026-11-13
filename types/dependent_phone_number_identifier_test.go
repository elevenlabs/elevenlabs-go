package types

import (
	"testing"
)

func TestDependentPhoneNumberIdentifierRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number_id": "value",
  "phone_number": "value",
  "label": "value",
  "provider": "twilio"
}`
	assertRoundTrip(t, wire, new(DependentPhoneNumberIdentifier))
}
