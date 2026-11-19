package types

import (
	"testing"
)

func TestTransferBranchInfoTrafficSplitRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "traffic_split",
  "branch_id": "value",
  "traffic_percentage": 9.99
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoTrafficSplit))
}

func TestTransferBranchInfoTrafficSplitRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "branch_reason": "traffic_split",
  "branch_id": "value",
  "traffic_percentage": 9.99
}`
	assertRoundTrip(t, wire, new(TransferBranchInfoTrafficSplit))
}
