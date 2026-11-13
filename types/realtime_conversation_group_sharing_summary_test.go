package types

import (
	"testing"
)

func TestRealtimeConversationGroupSharingSummaryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "people_count": 123,
  "shared_with_everyone": true,
  "everyone_role": "admin"
}`
	assertRoundTrip(t, wire, new(RealtimeConversationGroupSharingSummary))
}
