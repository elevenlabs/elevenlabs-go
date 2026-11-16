package types

import (
	"testing"
)

func TestConversationInitiationSourceInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "unknown",
  "version": "value"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationSourceInfo))
}

func TestConversationInitiationSourceInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "unknown",
  "version": "value"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationSourceInfo))
}

func TestConversationInitiationSourceInfoRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "unknown",
  "version": "value"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationSourceInfo))
}
