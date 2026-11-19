package types

import (
	"testing"
)

func TestAudioIsolationHistoryItemResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "title": "value",
  "created_at_unix": 123,
  "format": "value",
  "duration_seconds": 9.99,
  "download_url": "value",
  "icon_url": "value",
  "source_video_url": "value",
  "supports_video": true,
  "processing": true,
  "video_processing_failed": true,
  "preview_b64": "value"
}`
	assertRoundTrip(t, wire, new(AudioIsolationHistoryItemResponseModel))
}
