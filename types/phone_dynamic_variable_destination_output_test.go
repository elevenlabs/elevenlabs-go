package types

import (
	"testing"
)

func TestPhoneDynamicVariableDestinationOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "phone_dynamic_variable"
}`
	assertRoundTrip(t, wire, new(PhoneDynamicVariableDestinationOutput))
}
