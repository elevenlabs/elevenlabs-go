package types

import (
	"testing"
)

func TestCreateOrderParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_order"
}`
	assertRoundTrip(t, wire, new(CreateOrderParams))
}
