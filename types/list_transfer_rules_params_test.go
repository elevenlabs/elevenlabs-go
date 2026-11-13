package types

import (
	"testing"
)

func TestListTransferRulesParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_transfer_rules"
}`
	assertRoundTrip(t, wire, new(ListTransferRulesParams))
}
