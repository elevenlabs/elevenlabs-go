package types

import (
	"testing"
)

func TestDeleteTransferRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_transfer_rule"
}`
	assertRoundTrip(t, wire, new(DeleteTransferRuleParams))
}
