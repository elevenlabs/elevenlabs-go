package types

import (
	"testing"
)

func TestOptInSmsReminderParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "opt_in_sms_reminder"
}`
	assertRoundTrip(t, wire, new(OptInSmsReminderParams))
}
