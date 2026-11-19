package types

import (
	"testing"
)

func TestPodcastConversationModeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "conversation",
  "conversation": {
    "host_voice_id": "value",
    "guest_voice_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(PodcastConversationMode))
}
