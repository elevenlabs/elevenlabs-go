package types

import (
	"testing"
)

func TestTransferToAgentToolResultSuccessModelInputBranchInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "configured",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferToAgentToolResultSuccessModelInputBranchInfo))
}
