package types

import (
	"testing"
)

func TestChapterContentBlockInputModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sub_type": "p",
  "nodes": [
    {
      "type": "tts_node",
      "text": "value",
      "voice_id": "value"
    }
  ],
  "block_id": "value"
}`
	assertRoundTrip(t, wire, new(ChapterContentBlockInputModel))
}
