package types

import (
	"testing"
)

func TestPodcastBulletinModeDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "host_voice_id": "value"
}`
	assertRoundTrip(t, wire, new(PodcastBulletinModeData))
}
