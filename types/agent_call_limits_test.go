package types

import (
	"testing"
)

func TestAgentCallLimitsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_concurrency_limit": 123,
  "daily_limit": 123,
  "bursting_enabled": true
}`
	assertRoundTrip(t, wire, new(AgentCallLimits))
}

func TestAgentCallLimitsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_concurrency_limit": 123,
  "daily_limit": 123,
  "bursting_enabled": true
}`
	assertRoundTrip(t, wire, new(AgentCallLimits))
}
