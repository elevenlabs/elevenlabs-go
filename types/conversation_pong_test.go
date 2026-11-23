package types

import (
	"testing"
)

func TestConversationPongRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "pong",
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(ConversationPong))
}
