package types

import (
	"testing"
)

func TestRenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "render_type": "mp4",
  "normalize_volume": true
}`
	assertRoundTrip(t, wire, new(RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost))
}
