package types

import (
	"testing"
)

func TestCampaignSummaryCountsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "leads": 123,
  "terminal_leads": 123,
  "open_attention": 123,
  "delivery": {
    "sent": 123,
    "delivered": 123,
    "read": 123,
    "replied": 123,
    "failed": 123
  }
}`
	assertRoundTrip(t, wire, new(CampaignSummaryCounts))
}
