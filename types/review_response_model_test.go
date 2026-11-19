package types

import (
	"testing"
)

func TestReviewResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "review_status": "approved",
  "reviewed_at_unix": 123,
  "reviewed_by": "value",
  "reject_reasons": [
    "lacks_structure"
  ],
  "scores_breakdown": {
    "key": 123
  },
  "rejected_details": "value",
  "explanation": "value"
}`
	assertRoundTrip(t, wire, new(ReviewResponseModel))
}
