package types

import (
	"testing"
)

func TestPauseAllCampaignsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "paused_campaign_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(PauseAllCampaignsResponse))
}
