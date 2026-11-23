package types

import (
	"testing"
)

func TestAgentResponseCorrectionAgentResponseCorrectionEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "original_agent_response": "value",
  "corrected_agent_response": "value",
  "event_id": 123,
  "response_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentResponseCorrectionAgentResponseCorrectionEvent))
}
