package agents

import (
	"testing"
)

func TestBodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_branch_id": "value",
  "target_branch_id": "value",
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "triage_ticket_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost))
}

func TestBodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "requested_reviewer_user_ids": [
    "value"
  ],
  "close": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch))
}

func TestBodyReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "state": "approved",
  "comment": "value"
}`
	assertRoundTrip(t, wire, new(BodyReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost))
}

func TestBodyCommentOnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "body": "value"
}`
	assertRoundTrip(t, wire, new(BodyCommentOnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost))
}

func TestBodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "archive_source_branch": true,
  "force": true
}`
	assertRoundTrip(t, wire, new(BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost))
}
