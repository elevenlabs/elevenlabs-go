package types

import (
	"testing"
)

func TestTransferToAgentToolResultErrorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_agent_error",
  "status": "error",
  "from_agent": "value",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolResultErrorModel))
}

func TestTransferToAgentToolResultErrorModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "transfer_to_agent_error",
  "status": "error",
  "from_agent": "value",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolResultErrorModel))
}
