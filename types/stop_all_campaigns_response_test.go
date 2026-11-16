package types

import (
	"testing"
)

func TestStopAllCampaignsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stopped_campaign_ids": [
    "value"
  ],
  "stopping_campaign_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(StopAllCampaignsResponse))
}
