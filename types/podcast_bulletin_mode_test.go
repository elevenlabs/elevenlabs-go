package types

import (
	"testing"
)

func TestPodcastBulletinModeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "bulletin",
  "bulletin": {
    "host_voice_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(PodcastBulletinMode))
}
