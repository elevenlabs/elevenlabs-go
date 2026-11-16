package types

import (
	"testing"
)

func TestOptOutSmsReminderParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "opt_out_sms_reminder"
}`
	assertRoundTrip(t, wire, new(OptOutSmsReminderParams))
}
