package types

import (
	"testing"
)

func TestPodcastUrlSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "url",
  "url": "value"
}`
	assertRoundTrip(t, wire, new(PodcastUrlSource))
}
