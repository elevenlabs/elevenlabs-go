package types

import (
	"testing"
)

func TestChapterContentInputModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "blocks": [
    {
      "sub_type": "p",
      "nodes": [
        {
          "type": "tts_node",
          "text": "value",
          "voice_id": "value"
        }
      ],
      "block_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ChapterContentInputModel))
}
