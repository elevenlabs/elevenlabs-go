package types

import (
	"testing"
)

func TestTransferToAgentToolResultSuccessModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_agent_success",
  "status": "success",
  "from_agent": "value",
  "to_agent": "value",
  "to_node": "value",
  "condition": "value",
  "delay_ms": 123,
  "transfer_message": "value",
  "enable_transferred_agent_first_message": true,
  "branch_info": {
    "branch_reason": "configured",
    "branch_id": "value"
  },
  "preserve_client_tts_overrides": true,
  "preserve_voice_settings": true
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolResultSuccessModelInput))
}
