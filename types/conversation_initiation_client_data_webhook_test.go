package types

import (
	"testing"
)

func TestConversationInitiationClientDataWebhookRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "request_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataWebhook))
}

func TestConversationInitiationClientDataWebhookRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "request_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataWebhook))
}
