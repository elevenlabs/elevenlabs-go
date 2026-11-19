package types

import (
	"testing"
)

func TestAgentReasoningResponsePartClientEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "reasoning_response_part": {
    "text": "value",
    "type": "start",
    "event_id": 123
  },
  "type": "agent_reasoning_response_part"
}`
	assertRoundTrip(t, wire, new(AgentReasoningResponsePartClientEvent))
}
