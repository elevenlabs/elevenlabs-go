package types

import (
	"testing"
)

func TestBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "text",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem))
}
