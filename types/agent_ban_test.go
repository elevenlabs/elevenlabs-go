package types

import (
	"testing"
)

func TestAgentBanRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "at_unix": 123,
  "reason": "value",
  "reason_type": "safety"
}`
	assertRoundTrip(t, wire, new(AgentBan))
}
