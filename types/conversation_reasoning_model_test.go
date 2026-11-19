package types

import (
	"testing"
)

func TestConversationReasoningModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "summary": "value",
  "provider_redact": true
}`
	assertRoundTrip(t, wire, new(ConversationReasoningModel))
}

func TestConversationReasoningModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "summary": "value",
  "provider_redact": true
}`
	assertRoundTrip(t, wire, new(ConversationReasoningModel))
}

func TestConversationReasoningModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "summary": "value",
  "provider_redact": true
}`
	assertRoundTrip(t, wire, new(ConversationReasoningModel))
}
