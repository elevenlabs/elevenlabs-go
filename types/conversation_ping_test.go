package types

import (
	"testing"
)

func TestConversationPingRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "ping_event": {
    "event_id": 123,
    "ping_ms": 123
  },
  "type": "ping"
}`
	assertRoundTrip(t, wire, new(ConversationPing))
}
