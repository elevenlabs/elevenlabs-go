package types

import (
	"testing"
)

func TestTransferBranchInfoDefaultingToMainRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "defaulting_to_main",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoDefaultingToMain))
}

func TestTransferBranchInfoDefaultingToMainRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "defaulting_to_main",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoDefaultingToMain))
}
