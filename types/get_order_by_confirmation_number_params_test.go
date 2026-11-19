package types

import (
	"testing"
)

func TestGetOrderByConfirmationNumberParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_order_by_confirmation_number"
}`
	assertRoundTrip(t, wire, new(GetOrderByConfirmationNumberParams))
}
