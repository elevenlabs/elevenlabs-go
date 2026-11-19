package types

import (
	"testing"
)

func TestPhoneNumberAgentInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "agent_name": "value",
  "environment": "value",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(PhoneNumberAgentInfo))
}
