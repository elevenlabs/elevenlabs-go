package types

import (
	"testing"
)

func TestConversationConfigOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": 123
}`
	assertRoundTrip(t, wire, new(ConversationConfigOverride))
}

func TestConversationConfigOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": 123
}`
	assertRoundTrip(t, wire, new(ConversationConfigOverride))
}
