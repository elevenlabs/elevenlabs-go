package types

import (
	"testing"
)

func TestGenerationSourceContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "generation",
  "generation_id": "value",
  "prompt": "value",
  "model_id": "value",
  "model_provider": "value",
  "generation_session_id": "value",
  "session_iteration_id": "value",
  "model_parameters": {
    "key": {}
  },
  "extend_video": {
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
  }
}`
	assertRoundTrip(t, wire, new(GenerationSourceContext))
}
