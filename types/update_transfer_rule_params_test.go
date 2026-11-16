package types

import (
	"testing"
)

func TestUpdateTransferRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_transfer_rule",
  "post_dial_digits_enabled": true
}`
	assertRoundTrip(t, wire, new(UpdateTransferRuleParams))
}
