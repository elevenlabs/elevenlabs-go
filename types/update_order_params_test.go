package types

import (
	"testing"
)

func TestUpdateOrderParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_order"
}`
	assertRoundTrip(t, wire, new(UpdateOrderParams))
}
