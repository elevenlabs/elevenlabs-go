package types

import (
	"testing"
)

func TestMergeProposalReviewRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user_id": "value",
  "state": "approved",
  "submitted_at": 123,
  "comment": "value",
  "reviewer_role": "admin"
}`
	assertRoundTrip(t, wire, new(MergeProposalReview))
}
