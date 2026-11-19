package types

import (
	"testing"
)

func TestCreateTransferRuleParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_transfer_rule",
  "post_dial_digits_enabled": true
}`
	assertRoundTrip(t, wire, new(CreateTransferRuleParams))
}
