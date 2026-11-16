package types

import (
	"testing"
)

func TestStudioClipLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "chapter_id": "value",
  "clip_type": "video",
  "clip_id": "value",
  "block_id": "value",
  "preview_url": "value"
}`
	assertRoundTrip(t, wire, new(StudioClipLocator))
}
