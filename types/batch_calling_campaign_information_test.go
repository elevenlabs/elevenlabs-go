package types

import (
	"testing"
)

func TestBatchCallingCampaignInformationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "campaign_id": "value",
  "campaign_lead_id": "value",
  "cold_attempt": 123,
  "conversation_type": "initial"
}`
	assertRoundTrip(t, wire, new(BatchCallingCampaignInformation))
}
