package types

import (
	"testing"
)

func TestTicketMergeProposalLinkResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "merge_proposal_id": "value",
  "linked_by_user_id": "value",
  "linked_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(TicketMergeProposalLinkResponseModel))
}
