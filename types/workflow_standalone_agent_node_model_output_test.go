package types

import (
	"testing"
)

func TestWorkflowStandaloneAgentNodeModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "standalone_agent",
  "position": {
    "x": 9.99,
    "y": 9.99
  },
  "edge_order": [
    "value"
  ],
  "agent_id": "value",
  "branch_id": "value",
  "node_id": "value",
  "delay_ms": 123,
  "transfer_message": "value",
  "enable_transferred_agent_first_message": true,
  "preserve_client_tts_overrides": true,
  "preserve_voice_settings": true
}`
	assertRoundTrip(t, wire, new(WorkflowStandaloneAgentNodeModelOutput))
}
