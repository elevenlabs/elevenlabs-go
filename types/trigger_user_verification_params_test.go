package types

import (
	"testing"
)

func TestTriggerUserVerificationParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "trigger_user_verification"
}`
	assertRoundTrip(t, wire, new(TriggerUserVerificationParams))
}
