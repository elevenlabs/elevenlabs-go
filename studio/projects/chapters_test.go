package projects

import (
	"testing"
)

func TestBodyCreateChapterV1StudioProjectsProjectIdChaptersPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "from_url": "value"
}`
	assertRoundTrip(t, wire, new(BodyCreateChapterV1StudioProjectsProjectIdChaptersPost))
}

func TestBodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "content": {
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
  }
}`
	assertRoundTrip(t, wire, new(BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost))
}
