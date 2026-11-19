package types

import (
	"testing"
)

func TestSipUriDynamicVariableDestinationOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "destination": "value",
  "type": "sip_uri_dynamic_variable"
}`
	assertRoundTrip(t, wire, new(SipUriDynamicVariableDestinationOutput))
}
