package types

import (
	"testing"
)

func TestAgentResponseCorrectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_response_correction_event": {
    "original_agent_response": "value",
    "corrected_agent_response": "value",
    "event_id": 123,
    "response_id": "value"
  },
  "type": "agent_response_correction"
}`
	assertRoundTrip(t, wire, new(AgentResponseCorrection))
}
