package types

import (
	"testing"
)

func TestSelfSubagentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "execution_mode": "immediate",
  "timeout_secs": 123,
  "description": "value"
}`
	assertRoundTrip(t, wire, new(SelfSubagent))
}

func TestSelfSubagentRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "execution_mode": "immediate",
  "timeout_secs": 123,
  "description": "value"
}`
	assertRoundTrip(t, wire, new(SelfSubagent))
}
