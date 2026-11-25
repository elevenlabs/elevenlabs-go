package types

import (
	"testing"
)

func TestAgentResponseCompleteRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_response_complete_event": {
    "event_id": 123
  },
  "type": "agent_response_complete"
}`
	assertRoundTrip(t, wire, new(AgentResponseComplete))
}
