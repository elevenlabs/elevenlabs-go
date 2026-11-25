package types

import (
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideAsrRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "keywords": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataConversationConfigOverrideAsr))
}
