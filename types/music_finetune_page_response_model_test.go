package types

import (
	"testing"
)

func TestMusicFinetunePageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "finetunes": [
    {
      "id": "value",
      "name": "value",
      "tags": [
        "value"
      ],
      "primary_genre": "value",
      "model_id": "value",
      "created_at": "2024-01-01T00:00:00Z",
      "visibility": "private",
      "created_by": "self",
      "status": "pending",
      "training_progress": 9.99,
      "failure_reason": "audio_processing_failed"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(MusicFinetunePageResponseModel))
}
