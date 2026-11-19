package types

import (
	"testing"
)

func TestGetSpeechHistoryResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "history": [
    {
      "history_item_id": "value",
      "request_id": "value",
      "voice_id": "value",
      "model_id": "value",
      "voice_name": "value",
      "voice_category": "premade",
      "text": "value",
      "date_unix": 123,
      "character_count_change_from": 123,
      "character_count_change_to": 123,
      "content_type": "value",
      "state": {},
      "settings": {
        "key": {}
      },
      "share_link_id": "value",
      "source": "TTS",
      "alignments": {
        "alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        },
        "normalized_alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        }
      },
      "dialogue": [
        {
          "text": "value",
          "voice_id": "value",
          "voice_name": "value"
        }
      ],
      "output_format": "value"
    }
  ],
  "last_history_item_id": "value",
  "has_more": true,
  "scanned_until": 123
}`
	assertRoundTrip(t, wire, new(GetSpeechHistoryResponse))
}
