package types

import (
	"testing"
)

func TestUpdateBusinessInfoParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_business_info"
}`
	assertRoundTrip(t, wire, new(UpdateBusinessInfoParams))
}
