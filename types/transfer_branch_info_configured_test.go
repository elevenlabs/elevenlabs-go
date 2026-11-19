package types

import (
	"testing"
)

func TestTransferBranchInfoConfiguredRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "configured",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoConfigured))
}

func TestTransferBranchInfoConfiguredRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "configured",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoConfigured))
}
