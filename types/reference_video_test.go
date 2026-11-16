package types

import (
	"testing"
)

func TestReferenceVideoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "generation_id": "value",
  "content_asset_id": "value",
  "template_node_id": "value",
  "studio_clip": {
    "project_id": "value",
    "chapter_id": "value",
    "clip_type": "video",
    "clip_id": "value",
    "block_id": "value",
    "preview_url": "value"
  }
}`
	assertRoundTrip(t, wire, new(ReferenceVideo))
}
