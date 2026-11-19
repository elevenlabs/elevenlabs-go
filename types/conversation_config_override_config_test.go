package types

import (
	"testing"
)

func TestConversationConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": true
}`
	assertRoundTrip(t, wire, new(ConversationConfigOverrideConfig))
}
