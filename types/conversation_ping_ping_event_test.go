package types

import (
	"testing"
)

func TestConversationPingPingEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "event_id": 123,
  "ping_ms": 123
}`
	assertRoundTrip(t, wire, new(ConversationPingPingEvent))
}
