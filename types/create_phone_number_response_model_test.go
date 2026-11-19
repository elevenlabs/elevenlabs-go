package types

import (
	"testing"
)

func TestCreatePhoneNumberResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number_id": "value"
}`
	assertRoundTrip(t, wire, new(CreatePhoneNumberResponseModel))
}
