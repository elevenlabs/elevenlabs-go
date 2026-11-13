package outboundcampaigns

import (
	"testing"
)

func TestAttentionResolutionRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resolution": "value"
}`
	assertRoundTrip(t, wire, new(AttentionResolutionRequest))
}
