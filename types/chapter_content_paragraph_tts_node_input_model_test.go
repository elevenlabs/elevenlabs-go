package types

import (
	"testing"
)

func TestChapterContentParagraphTtsNodeInputModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "tts_node",
  "text": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ChapterContentParagraphTtsNodeInputModel))
}
