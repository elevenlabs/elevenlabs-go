package types

import (
	"testing"
)

func TestAgentReasoningResponsePartClientEventDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "type": "start",
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(AgentReasoningResponsePartClientEventData))
}
