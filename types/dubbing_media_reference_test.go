package types

import (
	"testing"
)

func TestDubbingMediaReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "src": "value",
  "content_type": "value",
  "bucket_name": "value",
  "random_path_slug": "value",
  "duration_secs": 9.99,
  "is_audio": true,
  "url": "value"
}`
	assertRoundTrip(t, wire, new(DubbingMediaReference))
}
