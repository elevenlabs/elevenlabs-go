package types

import (
	"testing"
)

func TestCreateAgentMergeProposalResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_merge_proposal_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentMergeProposalResponseModel))
}
