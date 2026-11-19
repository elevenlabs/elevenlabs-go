package types

import (
	"testing"
)

func TestRenderRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "version": 123,
  "language": "value",
  "type": "mp4",
  "media_ref": {
    "src": "value",
    "content_type": "value",
    "bucket_name": "value",
    "random_path_slug": "value",
    "duration_secs": 9.99,
    "is_audio": true,
    "url": "value"
  },
  "status": "complete"
}`
	assertRoundTrip(t, wire, new(Render))
}
