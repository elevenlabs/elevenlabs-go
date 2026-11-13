package types

import (
	"testing"
)

func TestBodyCreatePodcastV1StudioPodcastsPostModeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "conversation",
  "conversation": {
    "host_voice_id": "value",
    "guest_voice_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(BodyCreatePodcastV1StudioPodcastsPostMode))
}
