package types

import (
	"testing"
)

func TestTransferToAgentToolResultSuccessModelOutputBranchInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "configured",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolResultSuccessModelOutputBranchInfo))
}
