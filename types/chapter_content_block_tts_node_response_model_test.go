package types

import (
	"testing"
)

func TestChapterContentBlockTtsNodeResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "tts_node",
  "project_voice_ref_id": "value",
  "text": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ChapterContentBlockTtsNodeResponseModel))
}
