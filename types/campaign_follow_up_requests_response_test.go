package types

import (
	"testing"
)

func TestCampaignFollowUpRequestsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "callbacks": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  },
  "redirects": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  }
}`
	assertRoundTrip(t, wire, new(CampaignFollowUpRequestsResponse))
}
