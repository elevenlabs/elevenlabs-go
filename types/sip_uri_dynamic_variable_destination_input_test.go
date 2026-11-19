package types

import (
	"testing"
)

func TestSipUriDynamicVariableDestinationInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "sip_uri_dynamic_variable"
}`
	assertRoundTrip(t, wire, new(SipUriDynamicVariableDestinationInput))
}
