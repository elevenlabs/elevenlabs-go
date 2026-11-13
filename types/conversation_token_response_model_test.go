package types

import (
	"testing"
)

func TestConversationTokenResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "conversation_token": "value",
  "expiration_time_unix_secs": 123,
  "conversation_id": "value",
  "purpose": "signed_url",
  "token_requester_user_id": "value"
}`
	assertRoundTrip(t, wire, new(ConversationTokenResponseModel))
}
