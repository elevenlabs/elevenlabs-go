package types

import (
	"testing"
)

func TestSuppressLeadResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "campaign_id": "value",
  "lead_id": "value",
  "has_calls_in_flight": true
}`
	assertRoundTrip(t, wire, new(SuppressLeadResponse))
}
