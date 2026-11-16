package types

import (
	"testing"
)

func TestAgentFailureResponseExampleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "response": "value",
  "type": "failure"
}`
	assertRoundTrip(t, wire, new(AgentFailureResponseExample))
}

func TestAgentFailureResponseExampleRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "response": "value",
  "type": "failure"
}`
	assertRoundTrip(t, wire, new(AgentFailureResponseExample))
}
