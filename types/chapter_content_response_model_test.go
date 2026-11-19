package types

import (
	"testing"
)

func TestChapterContentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "blocks": [
    {
      "block_id": "value",
      "nodes": [
        {
          "type": "tts_node",
          "project_voice_ref_id": "value",
          "text": "value",
          "voice_id": "value"
        }
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(ChapterContentResponseModel))
}
