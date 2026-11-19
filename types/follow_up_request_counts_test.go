package types

import (
	"testing"
)

func TestFollowUpRequestCountsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "requested": 123,
  "requested_by_disposition": {
    "key": 123
  },
  "distinct_leads": 123
}`
	assertRoundTrip(t, wire, new(FollowUpRequestCounts))
}
