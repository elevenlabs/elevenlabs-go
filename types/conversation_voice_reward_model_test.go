package types

import (
	"testing"
)

func TestConversationVoiceRewardModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "reward_usd_cents": 9.99
}`
	assertRoundTrip(t, wire, new(ConversationVoiceRewardModel))
}
