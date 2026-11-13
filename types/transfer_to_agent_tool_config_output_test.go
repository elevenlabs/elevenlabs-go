package types

import (
	"testing"
)

func TestTransferToAgentToolConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "transfer_to_agent",
  "transfers": [
    {
      "agent_id": "value",
      "branch_id": "value",
      "node_id": "value",
      "condition": "value",
      "delay_ms": 123,
      "transfer_message": "value",
      "enable_transferred_agent_first_message": true,
      "is_workflow_node_transfer": true,
      "preserve_client_tts_overrides": true,
      "preserve_voice_settings": true
    }
  ]
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolConfigOutput))
}
