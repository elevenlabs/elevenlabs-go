package types

import (
	"testing"
)

func TestCancelOrderParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "cancel_order"
}`
	assertRoundTrip(t, wire, new(CancelOrderParams))
}
