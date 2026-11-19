package types

import (
	"testing"
)

func TestAgentSuccessfulResponseExampleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "response": "value",
  "type": "success"
}`
	assertRoundTrip(t, wire, new(AgentSuccessfulResponseExample))
}

func TestAgentSuccessfulResponseExampleRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "response": "value",
  "type": "success"
}`
	assertRoundTrip(t, wire, new(AgentSuccessfulResponseExample))
}
