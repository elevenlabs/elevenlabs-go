package types

import (
	"testing"
)

func TestPodcastConversationModeDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "host_voice_id": "value",
  "guest_voice_id": "value"
}`
	assertRoundTrip(t, wire, new(PodcastConversationModeData))
}
