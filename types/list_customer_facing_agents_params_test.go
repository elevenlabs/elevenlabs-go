package types

import (
	"testing"
)

func TestListCustomerFacingAgentsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_customer_facing_agents"
}`
	assertRoundTrip(t, wire, new(ListCustomerFacingAgentsParams))
}
