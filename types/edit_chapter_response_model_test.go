package types

import (
	"testing"
)

func TestEditChapterResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
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
    }
  }
}`
	assertRoundTrip(t, wire, new(EditChapterResponseModel))
}
