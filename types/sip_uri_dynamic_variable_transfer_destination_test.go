package types

import (
	"testing"
)

func TestSipUriDynamicVariableTransferDestinationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "sip_uri_dynamic_variable",
  "sip_uri": "value"
}`
	assertRoundTrip(t, wire, new(SipUriDynamicVariableTransferDestination))
}
