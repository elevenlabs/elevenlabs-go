package types

import (
	"testing"
)

func TestMergeProposalCommentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "user_id": "value",
  "body": "value",
  "created_at": 123
}`
	assertRoundTrip(t, wire, new(MergeProposalComment))
}
