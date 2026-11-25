package types

import (
	"testing"
)

func TestConversationInitiationClientDataSourceInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "unknown",
  "version": "value"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataSourceInfo))
}
